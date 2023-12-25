package api

import (
	"encoding/json"
	"errors"
	"io"
	"local/pokedexcli/internal/pokecache"
	"net/http"
	"time"
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
			return errors.New("invalid location")
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