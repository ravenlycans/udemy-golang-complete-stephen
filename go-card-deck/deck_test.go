package main

import "testing"

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
