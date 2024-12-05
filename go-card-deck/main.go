package main

import "fmt"

func main() {
	deck := newDeck()

	deck.print()
	fmt.Println("*************")
	deck.shuffle()
	deck.print()
}
