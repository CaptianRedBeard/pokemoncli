package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("you must provide a location pokemon name")
	}

	pokemonName := args[0]

	pokemon, exists := cfg.CaughtPokemon[pokemonName]
	if !exists {
		fmt.Print("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types: \n")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %v\n", pokeType.Type.Name)
	}

	return nil
}
