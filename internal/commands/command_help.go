package commands

import "fmt"

func commandHelp(cfg *Config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, cmd := range SelectCmd() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
