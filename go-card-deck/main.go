package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}

	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Printf("%d: %s\n", i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
