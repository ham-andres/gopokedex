package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config) error
}

type config struct {
	Next	*string
	Previous	*string

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
	}
	return commands
}

func cleanInput(text string) []string {
	var result []string
	result = strings.Fields(strings.ToLower(text))
	return result
}

// below main function which handles the project

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	command := getCommands()

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		word := scanner.Text()
		words := cleanInput(word)
		if len(words) == 0 {
			continue
		}
		cmd, ok := command[words[0]]
		if ok {
			err := cmd.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	command := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n")
	for key, value := range command {
		fmt.Printf("%s: %s \n", key, value.Description)
	}
	return nil
}
