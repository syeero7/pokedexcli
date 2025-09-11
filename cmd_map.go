package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(c *Config) error {
	if (*c).Next == nil {
		return fmt.Errorf("map unavailable")
	}

	config, err := getPokemonLocations(*c.Next)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon areas: %w", err)
	}

	(*c).Next = config.Next
	(*c).Previous = config.Previous

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	if (*c).Previous == nil {
		return fmt.Errorf("mapb unavailable")
	}

	config, err := getPokemonLocations(*c.Previous)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon areas: %w", err)
	}

	(*c).Next = config.Next
	(*c).Previous = config.Previous

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func getPokemonLocations(url string) (Config, error) {
	res, err := http.Get(url)
	if err != nil {
		return Config{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Config{}, err
	}

	data := Config{}
	if err := json.Unmarshal(body, &data); err != nil {
		return Config{}, err
	}

	return data, nil
}
