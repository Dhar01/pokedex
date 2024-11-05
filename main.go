package main

import (
	"fmt"
)

func main() {
	var answer string

	for {
		fmt.Println("pokedex > ")
		fmt.Scanf("%s", &answer)
		figure(answer)
	}
}

func figure(choice string) {
	if choice == "help" {
		fmt.Println("\nWelcome to the Pokedex!")
		fmt.Println("Usage:")

		fmt.Println("\nhelp: Displays a help message")
		fmt.Println("exit: Exit the Pokedex")
	}
}