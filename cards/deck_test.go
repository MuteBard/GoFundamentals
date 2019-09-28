package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	//Test if deck has 52 cards
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("\nExpected:\tdeck length of %v\nReceived:\tdeck length of %v", 52, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("\nExpected:\tfirst card to be %s\nReceived:\tfirst card was %s", "Ace of Spades", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("\nExpected:\tlast card to be %s\nReceived:\tlast card was %s", "King of Clubs", d[len(d)-1])
	}

}

func TestShuffle(t *testing.T) {
	//Test if deck is shuffled
	d1, d2 := newDeck(), newDeck()
	d2.shuffle()

	if d1.toString() == d2.toString() {
		t.Error("\nExpected:\tdeck 1 to be in different order from deck 2\nReceived:\tdeck 1 same order as deck 2")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//Test if the deck can saved
	fileName := "_decktesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)
	if len(loadedDeck) != 52 {
		t.Errorf("\nExpected:\tdeck length of %v\nReceived:\tdeck length of %v", 52, len(loadedDeck))
	}

	os.Remove(fileName)

}
