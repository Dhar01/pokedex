package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func figureCmd() map[string]cliCommand {
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
	answer := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		answer.Scan()

		words := cleanInput(answer.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		cmd, exists := figureCmd()[commandName]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown cmd")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, cmd := range figureCmd() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
