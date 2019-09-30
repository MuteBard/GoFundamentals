package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	eb.printGreeting()
	sb.printGreeting()
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//does things
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	//does things
	return "Hola!"
}
