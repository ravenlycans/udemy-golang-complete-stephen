package main

import "fmt"

func main() {
	hand, _ := newDeck().deal(10)

	handStringify := hand.toString()
	handByteSlice := hand.toByteSlice()

	fmt.Println(handStringify)
	fmt.Println(handByteSlice)

	recoveredHand := byteSliceToDeck(handByteSlice)

	fmt.Println(recoveredHand)
}
