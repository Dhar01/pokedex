package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		cmd, exists := selectCmd()[commandName]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("invalid command!")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
