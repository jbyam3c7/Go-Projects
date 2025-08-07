package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Sup Ya'll!\n")
	fmt.Println("Welcome to my newest Text Based RPG, this time written in Go\n")
	fmt.Println("In this game, we'll join our protagonist (hero seems a bit heavy handed) in the middle of a search, but they just cant remember what they're looking for\n")
	fmt.Println("But first, our hero has a name, unfortunately they can't remember it. What do you think it is?\n")

	//keyboard input
	//Declare a reader to collect keyboard inputs
	reader := bufio.NewReader(os.Stdin)
	var playerName string

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Had a problem champ, wanna try again?", err)
		return
	}
	userInput = strings.TrimSuffix(userInput, "\n")

	playerName = userInput

	fmt.Printf("Ah, you want to call them %v? You sure? (Y/N)\n", playerName)
	userInput, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Had a problem champ, wanna try again?\n", err)
		return
	}

	userInput = strings.TrimSuffix(userInput, "\n")
	userInput = strings.ToLower(userInput)

	if userInput == "y" {
		fmt.Printf("Alright, %v it is! Not my name...\n", playerName)
	} else {
		fmt.Println("Well, what do you want to name them then?\n")
		fmt.Printf("And pick a better one this time...\n")
	}

}
