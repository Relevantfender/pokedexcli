package main

import (
	"time"

	"github.com/Relevantfender/pokedexcli/internal/pokeapi"
	"github.com/Relevantfender/pokedexcli/internal/pokecache"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
		pokeCache:     &pokeCache,
	}
	startRepl(cfg)
}
