package pokeapi

import (
	"net/http"
	"time"

	"github.com/Dhar01/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

// func checkError[T any](result T, err error) (T, error) {
// 	if err != nil {
// 		return result, err
// 	}
// 	return result, nil
// }
