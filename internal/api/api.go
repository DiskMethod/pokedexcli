package api

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/DiskMethod/pokedexcli/internal/pokedex"
	"github.com/DiskMethod/pokedexcli/internal/responses"
)

var currentLocationAreasURL string = "https://pokeapi.co/api/v2/location-area/"
var previousLocationAreasURL *string

func CommandMap([]string) error {
	jsonData := responses.Response{}
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

	jsonData := responses.Response{}
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
		jsonData := responses.LocationResponse{}
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
		return errors.New("incorrect usage. Usage: explore [location-area]")
	} 
}

func CommandCatch(args []string) error {
	argsLength := len(args)
	switch {
	case argsLength == 0:
		return errors.New("you didn't specify a pokemon. Usage: catch [pokemon-name]")
	case argsLength == 1:
		_, err := pokedex.Get(args[0])
		if err == nil {
			return errors.New("you've already caught this pokemon")
		}
		pokemon := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", args[0])
		jsonData := responses.PokemonResponse{}
		err = fetchJSON(pokemon, &jsonData)
		if err != nil {
			return err
		}
		
		const C float64 = 1000.0
		probability := math.Min(100.0, C / (float64(jsonData.BaseExperience) + 1))
		caught := float64(rand.Intn(100)) < probability

		fmt.Printf("Throwing a Pokeball at %s", args[0])
		for i := 0; i < 3; i += 1 {
			time.Sleep(time.Second)
			fmt.Printf(".")
		}
		if !caught {
			fmt.Printf("\n%s escaped!\n", args[0])
			return nil
		}

		fmt.Printf("\n%s was caught!\n", args[0])
		pokedex.Add(args[0], jsonData)

		return nil
	default:
		return errors.New("incorrect usage. Usage: catch [pokemon-name]")
	}
}