package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		word := cleanInput(input)[0]
		fmt.Println("Your command was:", word)
	}
}
