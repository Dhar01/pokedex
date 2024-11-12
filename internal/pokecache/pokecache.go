package pokecache

import (
	"log"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

// adding new entry to the cache
func (c Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// check if an entry exists or not
func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if data, ok := c.cache[key]; ok {
		log.Print("cache found!")
		return data.val, true
	}

	return nil, false
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.entryChecker(time.Now().UTC(), interval)
	}

}

// check if an entry is older or not, delete or update
func (c Cache) entryChecker(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, value := range c.cache {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.cache, key)
		}
	}
}
