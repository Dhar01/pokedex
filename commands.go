package main

import "github.com/Dhar01/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeApiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func selectCmd() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "map of next 20 locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "map of previous 20 locations",
			callback:    commandMapB,
		},
	}
}
