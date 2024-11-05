package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
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
