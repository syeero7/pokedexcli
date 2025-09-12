package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/syeero7/pokedexcli/internal/pokecache"
)

func GetLocationList(endpoint *string, cache *pokecache.Cache) (LocationList, error) {
	url := fmt.Sprintf("%s/location-area", apiURL)
	if endpoint != nil {
		url = *endpoint
	}

	locations := LocationList{}
	if data, ok := cache.Get(url); ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationList{}, err
		}

		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationList{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationList{}, err
	}

	if err := json.Unmarshal(body, &locations); err != nil {
		return LocationList{}, err
	}

	cache.Add(url, body)
	return locations, nil
}

func GetFoundPokemon(location string, cache *pokecache.Cache) (FoundPokemon, error) {
	url := fmt.Sprintf("%s/location-area/%s", apiURL, location)
	pokemon := FoundPokemon{}
	if data, ok := cache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return FoundPokemon{}, err
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return FoundPokemon{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return FoundPokemon{}, err
	}

	if err := json.Unmarshal(body, &pokemon); err != nil {
		return FoundPokemon{}, err
	}

	cache.Add(url, body)
	return pokemon, nil

}
