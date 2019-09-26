package main
import("fmt")
func main(){
	book := "abcABC"

	bookLen := len(book)

	fmt.Printf("There are %v bytes/chars in %v \n", bookLen, book)

	//each character corresponds to value in ASCII / bytes in the ansi escape sequences
	fmt.Println(book[0])
	fmt.Println(book[1])
	fmt.Println(book[2])
	fmt.Println(book[3])
	fmt.Println(book[4])
	fmt.Println(book[5])

	//slicing for abd
	fmt.Println(book[0:3])
	//slicing for ABC
	fmt.Println(book[3:6])

	// also `` exist 
	tictactoe :=
`
xox
xox
oo-
`
	 fmt.Println(tictactoe)
}

// $ go run strings.go
// There are 6 bytes/chars in abcABC 
// 97
// 98
// 99
// 65
// 66
// 67