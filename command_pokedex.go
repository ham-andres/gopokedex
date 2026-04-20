package main	

import (
	"fmt"
)

 // args ...string meaning it doesnt have to accept args 0,1 or many,
 // whereas if args []string the caller has to pass.
 
func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")

	
	for _,pokemons := range cfg.PokemonCaught {
		fmt.Printf("	- %s\n",pokemons.Name)
	}
	return nil
}
