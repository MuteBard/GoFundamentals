package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_Cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
