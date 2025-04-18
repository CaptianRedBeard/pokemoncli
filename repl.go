package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cliCommand represents a command with a name, description, and a callback function.
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Command registry to hold command names and their corresponding functions.
var commandRegistry = map[string]cliCommand{}

// startRepl starts a Read-Eval-Print Loop (REPL) for the Pokedex application.
func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	// Register the help command
	helpCommand := func () error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("")
		fmt.Println("help: Displays a help message")
		fmt.Println("exit: Exit the Pokedex")
		return nil
	}

	// Populate the command registry
	commandRegistry["help"] = cliCommand{
		name:        "help",
		description: "Display help message",
		callback:    helpCommand,
	}

	// Register the exit command
	exitCommand := func() error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		return nil
	}

	// Populate the command registry
	commandRegistry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    exitCommand,
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		// Look up the command in the registry

		if command, found := commandRegistry[commandName]; found {
			if err := command.callback(); err != nil {
				fmt.Printf("Error while executing command '%s': %v\n", commandName, err)
			}
			// Check if the command is "exit" and return to break the loop
			if commandName == "exit" {
				return // Exit the startRepl function
			}
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
		}
	}
}

// cleanInput takes a string input, converts it to lowercase, trims whitespace, and splits it into a slice of words.
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
