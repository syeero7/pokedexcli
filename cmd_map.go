package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/syeero7/pokedexcli/internal/pokecache"
)

type pokemonLocations struct {
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(c *Config) error {
	if (*c).nextLocationURL == nil {
		return fmt.Errorf("map unavailable")
	}

	locations, err := getPokemonLocations(*c.nextLocationURL, c.pokecache)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon areas: %w", err)
	}

	(*c).nextLocationURL = locations.Next
	(*c).prevLocationURL = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	if (*c).prevLocationURL == nil {
		return fmt.Errorf("mapb unavailable")
	}

	locations, err := getPokemonLocations(*c.prevLocationURL, c.pokecache)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon areas: %w", err)
	}

	(*c).nextLocationURL = locations.Next
	(*c).prevLocationURL = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func getPokemonLocations(url string, cache *pokecache.Cache) (pokemonLocations, error) {
	data := pokemonLocations{}

	if cachedData, ok := (*cache).Get(url); ok {

		if err := json.Unmarshal(cachedData, &data); err != nil {
			return pokemonLocations{}, err
		}

		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return pokemonLocations{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonLocations{}, err
	}

	(*cache).Add(url, body)
	if err := json.Unmarshal(body, &data); err != nil {
		return pokemonLocations{}, err
	}

	return data, nil
}
