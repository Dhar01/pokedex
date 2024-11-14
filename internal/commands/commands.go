package commands

import "github.com/Dhar01/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config, ...string) error
}

type Config struct {
	PokeApiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	CaughtPokemon   map[string]pokeapi.Pokemon
}

func SelectCmd() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "map of next 20 locations",
			Callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "map of previous 20 locations",
			Callback:    commandMapB,
		},

		"explore": {
			name:        "explore <location-name>",
			description: "Explore a location",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempt to catch a pokemon!",
			Callback:    commandCatch,
		},
	}
}
