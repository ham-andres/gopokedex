package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ham-andres/gopokedex/internal/pokeapi"

)

type config struct {
	pokeapiClient 	pokeapi.Client 
	NextLocationsURL	*string
	PrevLocationsURL	*string

}


// below main function which handles the project

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	command := getCommands()

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		cmd, exists := command[words[0]]

		if exists {
			err := cmd.Callback(cfg, words[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue

		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}


func cleanInput(text string) []string {
	var result []string
	result = strings.Fields(strings.ToLower(text))
	return result
}



type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config, []string ) error
}

// had to keep inside the func as the commands map is global and when
// when i execute go run . it cause an error

func getCommands() map[string]cliCommand {
	var commands = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:					"map",
			Description: 	"Get the next page of locations",
			Callback:			commandMapf,
		},
		"mapb": {
			Name: 				"mapb",
			Description:	"Get the previous page of locations",
			Callback: 		commandMapb,
		},
		"explore": {
			Name: 				"explore",
			Description: 	"Get the pokemon list of the locations",
			Callback: 		commandExplore,
		},

	}
	return commands
}

