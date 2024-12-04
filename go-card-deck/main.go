package main

func main() {
	startingHand, _ := newDeck().deal(26)
	hand := deck{}

	hand, startingHand = startingHand.deal(1)

	hand.print()

	hand, startingHand = startingHand.deal(1)

	hand.print()
}
