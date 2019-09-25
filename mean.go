package main
import("fmt")

func main(){
	var x int
	var y int

	x = 4
	y = 12

	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	var mean int
	mean = (x + y) / 2
	fmt.Printf("mean = %v, type of %T\n", mean, mean)
}

// $ go run mean.go
// x = 4, type of int
// y = 12, type of int
// mean = 8, type of int
