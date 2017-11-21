package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 4*4 {
		t.Errorf("Expected deck length to be 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected the last card to be Four of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testingFileName := "_decktesting"
	os.Remove(testingFileName)

	d := newDeck()
	d.saveToFile(testingFileName)

	loadedDeck := newDeckFromFile(testingFileName)
	if len(loadedDeck) != 4*4 {
		t.Errorf("Expected 16 cards but got %v", len(loadedDeck))
	}

	os.Remove(testingFileName)
}
