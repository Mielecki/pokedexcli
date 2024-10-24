package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type config struct {
	Next *string // URL of the next 20 locations 
	Previous *string // URL of the previous 20 locations
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
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
	}
}

// converts the input to lowercase, splits it into words by whitespaces and returns the first word from the split input
func normalizeInput(input string) (output string) {
	output = strings.ToLower(input)
	outputList := strings.Fields(output)
	if len(outputList) > 0 {
		return outputList[0]
	}
	return output
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		commandName := normalizeInput(scanner.Text())
		
		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(config); err != nil {
			fmt.Println(err)
		}
	}
}