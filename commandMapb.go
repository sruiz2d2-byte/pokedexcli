package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	url := *cfg.previous

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

	cfg.next = response.Next
	cfg.previous = response.Previous

	return nil
}
