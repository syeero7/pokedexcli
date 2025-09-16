package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/syeero7/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *Config) error {
	if len(cfg.arguments) == 0 {
		return errors.New("missing required argument 'pokemon_name'")
	}

	pokemonName := cfg.arguments[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := pokeapi.GetPokemonInfo(pokemonName, cfg.cache)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon data: %w", err)
	}

	catchStr := "escaped"
	if pokemon.BaseExperience <= rand.IntN(pokemon.BaseExperience*2) {
		catchStr = "was caught"
		cfg.pokedex[pokemonName] = struct{}{}
	}
	fmt.Printf("%s %s!\n", pokemonName, catchStr)

	return nil
}
