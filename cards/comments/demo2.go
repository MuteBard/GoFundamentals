// package main

// import "fmt"

// func main() {
// 	cards := []string{"Ace of Diamonds", newCard()}
// 	//Append does not modify the existing slice, but returns an exisiting slice
// 	cards = append(cards, "Six of Spades")

// 	//How to iterate over a closed set or slice
// 	//i = index of this element in the array
// 	//card = Current card we're iterating over
// 	//range cards = Take the slice of cards and loop over it
// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}
// }

// func newCard() string {
// 	return "Five of Diamonds"
// }

// //Go has two basic data structures for handling list of records
// //Arrays - Fixed length list of things
// //Slice - An array that can grow or shrink, every element must be of the same type
