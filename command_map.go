package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config) error {
	locationsForward, err := cfg.pokeAPIClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsForward.Next
	cfg.previousLocationURL = locationsForward.Previous

	for _, location := range locationsForward.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationsBackward, err := cfg.pokeAPIClient.ListLocationAreas(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsBackward.Next
	cfg.previousLocationURL = locationsBackward.Previous

	for _, location := range locationsBackward.Results {
		fmt.Println(location.Name)
	}

	return nil
}
