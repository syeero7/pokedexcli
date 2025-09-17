package main

import (
	"errors"
	"fmt"

	"github.com/syeero7/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *Config) error {
	if len(cfg.arguments) == 0 {
		return errors.New("missing required argument 'pokemon_name'")
	}

	pokemonName := cfg.arguments[0]
	if _, ok := cfg.pokedex[pokemonName]; !ok {
		return fmt.Errorf("you have not captured '%s'", pokemonName)
	}

	pokemon, err := pokeapi.GetPokemonInfo(pokemonName, cfg.cache)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon data: %w", err)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
