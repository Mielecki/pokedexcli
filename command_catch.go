package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Mielecki/pokedexcli/internal/pokeapi"
)

func commandCatch(config *config, pokemon string) error {
	if pokemon == "" {
		return errors.New("the name of the pokemon cannot be empty")
	}
	pokemonInfo, err := pokeapi.GetPokemonInfo(pokemon, &config.cache)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemonInfo.Name + "...")
	if catch(pokemonInfo.BaseExperience) {
		config.pokemonsCaught[pokemonInfo.Name] = pokemonInfo
		fmt.Println(pokemonInfo.Name + " was caught!")
	} else {
		fmt.Println(pokemonInfo.Name + " escaped!")
	}
	return nil
}

func catch(baseExperience int) bool {
	res := rand.Intn(baseExperience)
	return res <= 50
}