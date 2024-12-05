package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func stringToDeck(s string) deck {
	return deck(strings.Split(s, ";"))
}

func byteSliceToDeck(b []byte) deck {
	return stringToDeck(string(b))
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return deck{}
	}

	return byteSliceToDeck(bs)
}

func randomizeDeckOrder(d deck) deck {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		j := r.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%d: %s\n", i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ";")
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, d.toByteSlice(), 0666)
}

func (d deck) shuffle() deck {
	randDeck := d

	for i := 1; i <= 10; i++ {
		randDeck = randomizeDeckOrder(randDeck)
	}

	return randDeck
}
