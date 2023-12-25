package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/DiskMethod/pokedexcli/internal/pokecache"
)

var cache *pokecache.Cache = pokecache.NewCache(5 * time.Second)

func fetchJSON(url string, v interface{}) error {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		if res.StatusCode == 404 {
			return errors.New("invalid argument")
		}
		defer res.Body.Close()
	
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Set(url, body)
	}

	return json.Unmarshal(body, v)
}