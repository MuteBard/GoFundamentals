package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_Cards")

	cards := newDeckFromFile("my_Cards")
	cards.print()
}
