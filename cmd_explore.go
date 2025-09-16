package main

import (
	"errors"
	"fmt"

	"github.com/syeero7/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *Config) error {
	if len(cfg.arguments) == 0 {
		return errors.New("missing required argument 'location'")
	}

	location := cfg.arguments[0]
	fmt.Printf("Exploring %s...\n", location)
	pokemon, err := pokeapi.GetFoundPokemon(location, cfg.cache)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon data: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, poke := range pokemon.PokemonEncounters {
		fmt.Printf(" - %s\n", poke.Pokemon.Name)
	}

	return nil
}
