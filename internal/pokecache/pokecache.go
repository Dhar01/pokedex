package pokecache

import (
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

	c.reapLoop(interval)

	return c
}

func (c Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if data, ok := c.cache[key]; ok {
		return data.val, true
	}

	return nil, false
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		c.entryChecker(time.Now().UTC(), interval)
	// 	}
	// }

	for range ticker.C {
		c.entryChecker(time.Now().UTC(), interval)
	}

}

func (c Cache) entryChecker(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	// code to add

}
