package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#43FFa1",
	}
	fmt.Println(colors)

	pokemon := make(map[string]string)
	pokemon["Metagross"] = "Steel/Psychic"
	pokemon["Banette"] = "Ghost"
	fmt.Println(pokemon)
	delete(pokemon, "banette")
	fmt.Println(pokemon)

	pokemonNumber := make(map[int]string)
	pokemonNumber[376] = "Metagross"
	pokemonNumber[609] = "Chandelure"
	pokemonNumber[423] = "Gastrodon"
	pokemonNumber[474] = "Porygon-Z"
	pokemonNumber[389] = "Torterra"

	printMap(pokemonNumber)

}

func printMap(p map[int]string) {
	for id, name := range p {
		fmt.Printf("[id = %v, name = %s]\n", id, name)
	}
}
