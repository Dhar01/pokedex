package main

import (
	"time"

	"github.com/Dhar01/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeAPIClient: pokeClient,
	}

	startREPL(cfg)
}
