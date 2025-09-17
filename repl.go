package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/syeero7/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	nextLocationURL *string
	prevLocationURL *string
	arguments       []string
	cache           *pokecache.Cache
	pokedex         map[string]struct{}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.Trim(lowered, " ")
	return strings.Fields(trimmed)
}

func startREPL(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cfg.arguments = words[1:]
		command, ok := getCommands()[words[0]]
		if !ok {
			fmt.Printf("Unknown command %s\n", words[0])
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help information",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all Pokémon found at the specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the specified Pokémon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect captured Pokémon",
			callback:    commandInspect,
		},
	}
}
