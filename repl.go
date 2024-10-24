package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Mielecki/pokedexcli/internal/pokeapi"
	"github.com/Mielecki/pokedexcli/internal/pokecache"
)


type config struct {
	Next *string // URL of the next 20 locations 
	Previous *string // URL of the previous 20 locations
	cache pokecache.Cache
	pokemonsCaught map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name string
	description string
	callback func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Displays the names of next 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of previous 20 location areas",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays information about pokemons found in the location passed as an argument",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Gives the opportunity to throw pokeball at the pokemon passed as an argument and catch it",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Displays information about the caught pokemon passed as an argument",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays a list of all the names of the pokemon the user has caught",
			callback: commandPokedex,
		},
	}
}

type ParsedInput struct {
	Command string
	Argument string
}

// converts the input to lowercase, splits it into words by whitespaces and returns the first word from the split input
func parseInput(input string) (output ParsedInput) {
	input = strings.ToLower(input)
	inputList := strings.Fields(input)
	switch len(inputList) {
	case 0:
	case 1:
		output.Command = inputList[0]
	default:
		output.Command = inputList[0]
		output.Argument = inputList[1]
	}
	return output
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config.cache = pokecache.NewCache(5 * time.Minute)
	
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		parsedInput := parseInput(scanner.Text())
		
		command, exists := commands[parsedInput.Command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(config, parsedInput.Argument); err != nil {
			fmt.Println(err)
		}
	}
}