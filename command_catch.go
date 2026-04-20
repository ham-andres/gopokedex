package main 

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No Name Given")
	}
	pokemon_name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)
	pokemonInfo, err := cfg.pokeapiClient.GetPokemon(pokemon_name)
	if err != nil {
		return err
	}
	

	threshold := rand.Intn(300)
	
	if pokemonInfo.BaseExperience < threshold {
		
		cfg.PokemonCaught[pokemon_name] = pokemonInfo
		fmt.Printf("%s was caught! \n",pokemon_name)
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}	
	fmt.Printf("%s escaped!\n",pokemon_name)
	return nil
			
} 
