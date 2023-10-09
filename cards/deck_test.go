package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Expected deck length of 24, but got %v", len(d))
	}
	if d[len(d)-1] != "Six of Diamonds" {
		t.Errorf("Expected Six of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 24 {
		t.Errorf("Expected deck length of 24, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
