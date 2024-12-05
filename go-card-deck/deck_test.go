package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(deck))
	}
}
