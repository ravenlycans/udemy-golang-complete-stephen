package main

import "fmt"

func main() {
	hand, _ := newDeck().deal(10)

	handStringify := hand.toString()

	fmt.Println(handStringify)

	recoveredHand := stringToDeck(handStringify)

	fmt.Println(recoveredHand)
}
