package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	pokedex, err := cfg.pokeapiClient.ListPokemonsPokedex()

	if err != nil {
		return fmt.Errorf("error during getting the pokedex from cache")
	}

	fmt.Println("Your Pokedex:")
	for _, value := range pokedex {
		fmt.Printf("	- %s\n", value.Name)
	}
	return nil
}
