package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of Spades but got %v", d[0])
	}

	k := len(d) - 1

	if d[k] != "Four of Clubs" {
		t.Errorf("Expected last card as Four of Clubs but got %v", d[k])
	}

}

func TestFileCD(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_deckTesting")
}
