package main

import "fmt"

func commandPokedex(config *config, argument string) error {
	fmt.Println("Your pokedex: ")
	for k, _ := range config.pokemonsCaught {
		fmt.Println(" - " + k)
	}

	return nil
}