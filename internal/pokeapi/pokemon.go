package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/syeero7/pokedexcli/internal/pokecache"
)

func GetPokemonInfo(pokemonName string, cache *pokecache.Cache) (Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", apiURL, pokemonName)
	pokemon := Pokemon{}
	if data, ok := cache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	cache.Add(url, body)
	return pokemon, nil
}
