package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *Config) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("pokedex is empty. go catch some pokémon")
	}

	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
