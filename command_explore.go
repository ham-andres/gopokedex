package main 

import (
	"fmt"
)

// explore function
func commandExplore(cfg *config, args []string ) error {
		if len(args) == 0 {
			return fmt.Errorf("No Area Name Given")
		}
		area := args[0]
		pokemonList, err := cfg.pokeapiClient.LocationGet(area)
		if err != nil {
			return err
		}
		
		fmt.Printf("Exploring %s...\n", area)
		fmt.Println("Found Pokemon:")
		for _,pokemons := range pokemonList.PokemonEncounters {
			fmt.Printf(" - %s\n", pokemons.Pokemon.Name)
		}

		return nil
		

	} 
