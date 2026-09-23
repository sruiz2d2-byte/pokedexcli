package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExploreResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonReference `json:"pokemon"`
}

type PokemonReference struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area")
	}

	areaName := args[0]

	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	data, ok := cfg.cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, data)
	}

	var response ExploreResponse

	err := json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
