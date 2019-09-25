//introduction.go
package main

//Go is a very simple language, and by itself, does not have many built-in functions. Most of the code you'll be using will be in packages.
//The fmt package contains functions for formatted printing.
import (
	"fmt"
)
//We define a function with the func keyword.
//unlike java, you do not need to place a semi-colon at the end of the line
func main(){
	fmt.Println("Hello world")
}
//execute by entering go run introduction.go