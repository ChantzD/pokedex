package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name	string
	description	string
	callback	func(*config) error
}

type config struct {
	Next	string
	Previous string
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	defConfig := config{
		Next: "https://pokeapi.co/api/v2/location-area/?offset=0",
		Previous: "https://pokeapi.co/api/v2/location-area/?offset=0",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		args := cleanInput(input)
		if len(args) == 0{
			continue
		}

		commandName := args[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&defConfig)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Print the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Print the previous 20 locations",
			callback:    commandMapb,
		},
	}
}
