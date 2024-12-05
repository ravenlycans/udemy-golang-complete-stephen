package main

import "fmt"

func main() {

	loadedHand := newDeckFromFile("myDeck")

	if len(loadedHand) < 1 {
		fmt.Println("No cards where loaded")
	} else {
		fmt.Printf("Loaded hand: %s", loadedHand)
	}
}
