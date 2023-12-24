package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"local/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var currentLocationAreasURL string = "https://pokeapi.co/api/v2/location-area/"
var previousLocationAreasURL *string
var cache *pokecache.Cache = pokecache.NewCache(5 * time.Second)

func fetchJSON(url string, v interface{}) error {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
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

func CommandMap([]string) error {
	jsonData := Response{}
	err := fetchJSON(currentLocationAreasURL, &jsonData)
	if err != nil {
		return err
	}

	currentLocationAreasURL = jsonData.Next
	previousLocationAreasURL = jsonData.Previous

	for _, location := range jsonData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func CommandMapb([]string) error {
	if previousLocationAreasURL == nil {
		return errors.New("There are no previous map locations")
	}

	jsonData := Response{}
	err := fetchJSON(*previousLocationAreasURL, &jsonData)
	if err != nil {
		return err
	}

	currentLocationAreasURL = *previousLocationAreasURL
	previousLocationAreasURL = jsonData.Previous

	for _, location := range jsonData.Results {
		fmt.Println(location.Name)
	}

	return nil
}