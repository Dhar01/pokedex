package main

import (
	"time"

	"github.com/Dhar01/pokedexcli/internal/commands"
	"github.com/Dhar01/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute * 5)
	cfg := &commands.Config{
		PokeApiClient: pokeClient,
	}

	startREPL(cfg)
}
