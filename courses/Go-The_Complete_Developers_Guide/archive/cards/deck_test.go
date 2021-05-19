package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but instead got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but instead got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testfile := "/tmp/_decktesting"

	os.Remove(testfile)

	deck := newDeck()
	deck.saveToFile(testfile)

	loadedDeck := newDeckFromFile(testfile)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(loadedDeck))
	}

	os.Remove(testfile)
}
