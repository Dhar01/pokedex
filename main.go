package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func figureCommand() map[string]cliCommand {
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
	}
}

func main() {
	// var answer string
	answer := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		// fmt.Scan(&answer)
		// figure(answer)
		answer.Scan()
		// fmt.Println(answer.Text())

		command, exists := figureCommand()[answer.Text()]
		if exists {
			err := command.callback()
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

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, command := range figureCommand() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
	// fmt.Println("commandHelp")
	// return fmt.Errorf("this is commandHelp")

}

func commandExit() error {
	os.Exit(0)
	return nil

	// fmt.Println("commandExit")
	// return fmt.Errorf("this is commandExit")
}

// func figure(choice string) {
// 	if choice == "help" {
// 		fmt.Println("\nWelcome to the Pokedex!")
// 		fmt.Println("Usage:")
// 		fmt.Println("\nhelp: Displays a help message")
// 		fmt.Println("exit: Exit the Pokedex")
// 	}
// }
