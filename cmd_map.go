package main

import (
	"fmt"

	"github.com/syeero7/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	locations, err := pokeapi.GetLocationList(cfg.nextLocationURL, cfg.cache)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon areas: %w", err)
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.prevLocationURL == nil {
		return fmt.Errorf("mapb unavailable")
	}
	locations, err := pokeapi.GetLocationList(cfg.prevLocationURL, cfg.cache)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon areas: %w", err)
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}
