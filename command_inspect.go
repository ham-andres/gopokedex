package main	

import (
	"fmt"
)

func commandInspect( cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No Name Given")
	}
	pokemon_name := args[0]



	if val, ok :=cfg.PokemonCaught[pokemon_name]; ok {
		fmt.Printf("Name: %s \nHeight: %v \nWeight: %v \n",val.Name,val.Height,val.Weight)
		fmt.Println("Stats:")
		for _, stat := range val.Stats {
    	fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range val.Types {
			fmt.Println("	-",t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
} 
