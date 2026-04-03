package main

import ("strings"
	"fmt"
	"bufio"
	"os"
	)

func cleanInput(text string) []string {
	var result []string
	result = strings.Fields(strings.ToLower(text))
	return result
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")	
		scanner.Scan()
		word := scanner.Text()
		words := cleanInput(word)
		fmt.Printf("Your command was: %v \n",words[0])
	}
}
