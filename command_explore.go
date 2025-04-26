package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	// Check if any args were sent
	if len(args) == 0 {
		return errors.New("you must provide a location area name")
	}

	// Pull the first argument for the area to explore
	locationAreaName := args[0]

	// Debugging
	fmt.Printf("Exploring %s...\n", locationAreaName)

	// Using the API, pull the location area response
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	
	return nil
}
	
