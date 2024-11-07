package main

import "fmt"

func commandMapF(cfg *config) error {
	locationsForward, err := cfg.pokeAPIClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsForward.Next
	cfg.previousLocationURL = locationsForward.Previous

	for _, loc := range locationsForward.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	locationsBackward, err := cfg.pokeAPIClient.ListLocationAreas(cfg.previousLocationURL)
	if err != nil {
		return err
	}
}
