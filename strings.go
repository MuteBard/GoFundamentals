package main
import("fmt")
func main(){
	book := "this is book"

	bookLen := len(book)

	fmt.Printf("the length of the book is %v chars \n", bookLen)
}

// $ go run strings.go
// the length of the book is 12 chars 