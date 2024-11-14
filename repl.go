package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Dhar01/pokedexcli/internal/commands"
)

func startREPL(cfg *commands.Config) {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, ok := commands.SelectCmd()[commandName]
		if ok {
			err := cmd.Callback(cfg, args...)
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
