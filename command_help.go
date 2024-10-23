package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range getCommands() {
		fmt.Println(v.name + ": " + v.description)
	}
	fmt.Println()
	return nil
}