package main

import "github.com/Mielecki/pokedexcli/internal/pokeapi"

func main() {
	config := &config{}
	config.pokemonsCaught = map[string]pokeapi.Pokemon{}
	startRepl(config)
}