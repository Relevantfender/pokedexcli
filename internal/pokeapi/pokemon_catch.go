package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonStats(pokemon_name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon_name
	var pokemon Pokemon
	data, ok := c.cache.Get(pokemon_name)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error while creating a request: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error while recieving a response: %w", err)
		}

		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)

		if err != nil {
			return Pokemon{}, fmt.Errorf("error while reading the response: %w", err)
		}
		c.cache.Add(pokemon_name, data)

	}

	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error while unmarshaling the response: %w", err)
	}

	return pokemon, nil

}

func (c *Client) AddPokedex(pokemon_name string, pokemon_stats Pokemon) {
	if c.pokedex == nil {
		c.pokedex = make(map[string]Pokemon)
	}
	c.pokedex[pokemon_name] = pokemon_stats

}
func (c *Client) ListPokemonsPokedex() (map[string]Pokemon, error) {
	if c.pokedex == nil {
		return map[string]Pokemon{}, errors.New("no pokemon caught")
	}

	return c.pokedex, nil
}

func (c *Client) GetPokedexPokemonStats(pokemon_name string) (stats []pokemonStat, types []pokemonType) {
	pokemon, ok := c.pokedex[pokemon_name]

	if !ok {
		return nil, nil
	}
	statsList := pokemon.Stats

	typesList := pokemon.Types

	return statsList, typesList

}
