package main

import "fmt"

//Create a new type of 'deck'
//which is a slice of strings
//this new deck type that we are creating borrows all the behavior of a slice of type string
type deck []string

//() after func is known as a reciever and the function as a whole is a reciever on a function
//Any variable of type "deck" now gets access to the print method
// - d  Instance of the deck variable we're working with is available in the function as a variable called 'd'
// - deck  Every variable of type deck can call this function on itself
// by convention we indicate the reciever by a 1 or 2 letter abbrievation that matches the type of the reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}