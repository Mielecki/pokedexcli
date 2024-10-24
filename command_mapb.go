package main

import (
	"fmt"

	"github.com/Mielecki/pokedexcli/internal/pokeapi"
)

func commandMapb(config *config) error {
	res, err := pokeapi.GetLocationAreas(config.Previous)
	if err != nil {
		fmt.Println(err)
	}

	config.Next = res.Next
	config.Previous = res.Previous

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	return nil
}