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
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
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

// func commandHelp() error {
// 	fmt.Println("\nWelcome to the Pokedex!")
// 	fmt.Printf("Usage:\n\n")
// 	for _, cmd := range figureCmd() {
// 		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
// 	}
// 	fmt.Println()
// 	return nil
// }

// func commandExit() error {
// 	os.Exit(0)
// 	return nil
// }

// func commandMap() error {
// 	return nil
// }

// func commandMapB() error {
// 	return nil
// }
