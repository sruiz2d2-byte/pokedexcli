package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	// Si ya existe una página siguiente, usar esa URL
	if cfg.next != nil {
		url = *cfg.next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var response LocationAreasResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	// Guardar las URLs para futuras llamadas
	cfg.next = response.Next
	cfg.previous = response.Previous

	return nil
}
