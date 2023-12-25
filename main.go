package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DiskMethod/pokedexcli/internal/api"
	"github.com/DiskMethod/pokedexcli/internal/pokedex"
)

type cliCommand struct {
	name string
	description string
	callback func([]string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Subsequent calls display the next 20 locations.",
			callback: api.CommandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations.",
			callback: api.CommandMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays the pokemon within a location",
			callback: api.CommandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempts to catch pokemon",
			callback: api.CommandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspects a pokemon in your pokedex",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays all the pokemon in your pokedex",
			callback: commandPokedex,
		},
	}
}

func commandHelp([]string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit([]string) error {
	os.Exit(0)
	return nil
}

func commandInspect(args []string) error {
	pokemon, err := pokedex.Get(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex([]string) error {
	pokemonList, err := pokedex.List()
	if err != nil {
		return err
	}
	
	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokemonList {
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}

func main() {
	pokedex.NewPokedex()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		inputFields := strings.Fields(input)
		if len(inputFields) == 0 {
			continue
		}

		commandName := inputFields[0]
		args := inputFields[1:]

		cmd, ok := commands[commandName]
		if ok {
			err := cmd.callback(args)
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			fmt.Println("Invalid Command. Type \"help\" for more details")
		}
	}
}