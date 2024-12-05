package main

import (
	"errors"
	"os"
	"slices"
	"testing"
)

func TestNewDeckCount(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(deck))
	}
}

func TestNewDeckFirstCard(t *testing.T) {
	deck := newDeck()

	if deck[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs, got %s", deck[0])
	}
}

func TestNewDeckLastCard(t *testing.T) {
	deck := newDeck()

	if deck[51] != "King of Spades" {
		t.Errorf("Expected last card to be King of Spades, got %s", deck[51])
	}
}

func TestDeckShuffle(t *testing.T) {
	deck := newDeck()
	deck.shuffle()

	deckOrig := newDeck()

	if deckOrig.toString() == deck.toString() {
		t.Errorf("Expected deck to be shuffled, got unshuffled deck string %s", deck.toString())
	}
}

func TestDeckSaveToFile(t *testing.T) {
	_ = os.Remove("test.txt")

	deck := newDeck()
	err := deck.saveToFile("test.txt")

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	_, err = os.ReadFile("test.txt")

	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected file to be created, got %s", err)
	}

	_ = os.Remove("test.txt")
}

func TestNewDeckFromFile(t *testing.T) {
	_ = os.Remove("test.txt")

	deck := newDeck()
	err := deck.saveToFile("test.txt")

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	deckFromFile := newDeckFromFile("test.txt")
	if deck.toString() != deckFromFile.toString() {
		t.Errorf("Expected deck to be equal to deck from file, got %s", deck.toString())
	}

	_ = os.Remove("test.txt")
}

func TestDeckDeal(t *testing.T) {
	deck := newDeck()
	hand, remainingDeck := deck.deal(5)

	if len(hand) != 5 {
		t.Errorf("Expected 5 cards in hand, got %d", len(hand))
	}

	if len(remainingDeck) != 47 {
		t.Errorf("Expected 47 cards in remaining deck, got %d", len(remainingDeck))
	}

	if slices.Contains(remainingDeck, hand[0]) {
		t.Errorf("Expected hand[0] to be removed from remaining deck, got %s", remainingDeck)
	}
}
