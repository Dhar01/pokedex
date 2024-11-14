package commands

import "errors"

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	// pokemon, err :=

	return nil
}
