package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Thank you for using the Pokedex! Bye!")
	os.Exit(0)
	return nil
}
