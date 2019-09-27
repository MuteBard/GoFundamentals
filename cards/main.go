package main

func main() {
	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

//go run main.go deck.go
