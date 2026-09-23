package main

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int       `json:"base_stat"`
	Stat     NamedStat `json:"stat"`
}

type NamedStat struct {
	Name string `json:"name"`
}

// /////////
type PokemonType struct {
	Type NamedType `json:"type"`
}

type NamedType struct {
	Name string `json:"name"`
}
