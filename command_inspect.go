package main

import "fmt"

func commandInspect(config *config, pokemon string) error {
	pokemonInfo, ok := config.pokemonsCaught[pokemon]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name: " + pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats: ")
	for _, v := range pokemonInfo.Stats {
		fmt.Printf(" - %s: %d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types: ")
	for _, v := range pokemonInfo.Types {
		fmt.Printf(" - %s\n", v.Type.Name)
	}
	return nil
}