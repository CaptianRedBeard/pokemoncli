package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("you must provide a location pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchChance := rand.Intn(101)

	if catchChance-pokemon.BaseExperience < 0 {
		fmt.Print("The Pokémon broke free from the Poké Ball! Try again!\n")
		return nil
	}

	message := "Congratulations! You caught a " + pokemonName + "!"
	cfg.CaughtPokemon[pokemonName] = pokemon

	//It's been added to your Pokédex!"

	fmt.Printf("%s\n", message)

	return nil
}
