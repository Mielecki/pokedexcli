package main

import (
	"fmt"

	"github.com/Mielecki/pokedexcli/internal/pokeapi"
)

func commandExplore(config *config, location string) error {
	res, err := pokeapi.GetLocationAreaInfo(location, &config.cache)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + res.Location.Name + "...")
	for _, v := range res.PokemonEncounters {
		fmt.Println(" - " + v.Pokemon.Name)
	}

	return nil
}