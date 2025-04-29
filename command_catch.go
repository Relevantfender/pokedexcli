package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("insert the name of the pokemon")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonStats, err := cfg.pokeapiClient.GetPokemonStats(pokemonName)

	if err != nil {
		return err
	}

	isCaught := caughtCalculation(pokemonStats.BaseExperience)

	if !isCaught {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.pokeapiClient.AddPokedex(pokemonName, pokemonStats)
	return nil
}

func caughtCalculation(baseExp int) bool {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := rng.Intn(100) + 1

	catchProbability := 100 - (baseExp / 4)

	if catchProbability < 5 {
		catchProbability = 5
	} else if catchProbability > 90 {
		catchProbability = 90
	}

	if randomNum <= catchProbability {
		return true
	} else {
		return false
	}
}
