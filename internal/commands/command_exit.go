package commands

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Thank you for using the Pokedex! Bye!")
	os.Exit(0)
	return nil
}
