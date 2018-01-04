package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "The Ace of Hearts" {
		t.Errorf("Expected deck must The Ace of Hearts, but got %v", d[1])
	}

	if d[len(d)-1] != "The King of Club" {
		t.Errorf("Expected deck must The King of Club, but got %v", d[len(d)-1])
	}
}

func TestReadFile(t *testing.T) {
	os.Remove("file/_decktest.txt")

	d := newDeck()
	d.createToFile("file/_decktest.txt")

	fileRead := readFile("file/_decktest.txt")
	if len(fileRead) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	os.Remove("file/_decktest.txt")
}
