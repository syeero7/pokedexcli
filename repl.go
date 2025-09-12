package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

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
	pokecache       *pokecache.Cache
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.Trim(lowered, " ")
	return strings.Fields(trimmed)
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	locationURL := "https://pokeapi.co/api/v2/location-area/"
	config := Config{
		nextLocationURL: &locationURL,
		prevLocationURL: nil,
		pokecache:       pokecache.NewCache(5 * time.Second),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, ok := getCommands()[words[0]]
		if !ok {
			fmt.Printf("Unknown command %s\n", words[0])
			continue
		}

		err := command.callback(&config)
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
	}
}
