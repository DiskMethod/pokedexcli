package api

import (
	"errors"
	"fmt"
)

var currentLocationAreasURL string = "https://pokeapi.co/api/v2/location-area/"
var previousLocationAreasURL *string

func CommandMap([]string) error {
	jsonData := response{}
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
		return errors.New("there are no previous map locations")
	}

	jsonData := response{}
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

func CommandExplore(args []string) error {
	argsLength := len(args)
	switch {
	case argsLength == 0:
		return errors.New("you didn't specify a location. Usage: explore [location-area]")
	case argsLength == 1:
		fmt.Printf("Exploring %s...\n", args[0])
		locationArea := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", args[0])
		jsonData := locationResponse{}
		err := fetchJSON(locationArea, &jsonData)
		if err != nil {
			return err
		}

		fmt.Println("Found Pokemon:")
		for _, encounter := range jsonData.PokemonEncounters {
			fmt.Println("-", encounter.Pokemon.Name)
		}

		return nil
	default:
		return errors.New("incorrect usage. Example: explore [location-area]")
	} 
}