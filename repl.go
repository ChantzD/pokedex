package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		word := cleanInput(input)[0]
		fmt.Println("Your command was:", word)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
