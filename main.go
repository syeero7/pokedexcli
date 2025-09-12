package main

import (
	"time"

	"github.com/syeero7/pokedexcli/internal/pokecache"
)

func main() {
	config := Config{cache: pokecache.NewCache(5 * time.Minute)}
	startREPL(&config)
}
