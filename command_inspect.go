package main

import (
	"errors"
	"fmt"
)

// "errors"
// "fmt"

func commandInspect(cfg *config, args ...string) error {
	pokemon_name := args[0]

	stats, types := cfg.pokeapiClient.GetPokedexPokemonStats(pokemon_name)

	if stats == nil && types == nil {
		return errors.New("that pokemon isn't caught")
	}

	fmt.Println("Stats:")
	for _, stat := range stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range types {
		fmt.Printf("	-%s\n", t.Type.Name)
	}

	return nil
}
