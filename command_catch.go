package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a Pokemon name")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	data, ok := cfg.cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Pokemon %s not found", pokemonName)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, data)
	}

	var pokemon Pokemon

	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return err
	}

	chance := rand.Intn(600)

	if chance > pokemon.BaseExperience {
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
