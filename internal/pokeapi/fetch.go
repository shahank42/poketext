package pokeapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/shahank42/poketext/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Second)

func getWithCache(url string) ([]byte, bool, error) {
	data, foundCache := cache.Get(url)
	if !foundCache {
		res, err := http.Get(url)
		if err != nil {
			return nil, foundCache, err
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return nil, foundCache, errors.New(fmt.Sprintf("Response failed with status code %d", res.StatusCode))
		}
		if err != nil {
			return nil, foundCache, err
		}

		data = body
		cache.Add(url, data)
	}

	return data, foundCache, nil
}
