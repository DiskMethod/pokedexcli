package pokedex

import (
	"errors"
	"local/pokedexcli/internal/responses"
	"sync"
)

var (
    pdex map[string]responses.PokemonResponse
    once sync.Once
)

func NewPokedex() (map[string]responses.PokemonResponse, error) {
    once.Do(func() {
        pdex = make(map[string]responses.PokemonResponse)
    })
    if pdex == nil {
        return nil, errors.New("failed to create pokedex")
    }
    return pdex, nil
}

func Get(name string) (responses.PokemonResponse, error) {
    if pdex == nil {
        return responses.PokemonResponse{}, errors.New("pokedex is not created")
    }

    val, ok := pdex[name]
    if !ok {
        return responses.PokemonResponse{}, errors.New("pokemon not found")
    }

    return val, nil
}

func Add(name string, pokemon responses.PokemonResponse) error {
    if pdex == nil {
        return errors.New("pokedex is not created")
    }
    pdex[name] = pokemon

    return nil
}

func List() ([]string, error) {
	if pdex == nil {
		return nil, errors.New("pokedex is not created")
	}

	keys := []string{}
	for key := range pdex {
		keys = append(keys, key)
	}

	if len(keys) == 0 {
		return nil, errors.New("your pokedex is empty")
	}

	return keys, nil
}