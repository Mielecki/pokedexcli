package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
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
	}
}

// converts the input to lowercase, splits it into words by whitespaces and returns the first word from the split input
func normalizeInput(input string) (output string) {
	output = strings.ToLower(input)
	output = strings.Fields(output)[0]
	return output
}

func startRepl() {
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
		
		if err := command.callback(); err != nil {
			fmt.Println(err)
		}
	}
}