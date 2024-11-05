package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	nextLocationURL     string
	previousLocationURL string
}

func figureCmd() map[string]cliCommand {
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "map of previous 20 locations",
			callback:    commandMapB,
		},
	}
}
