package main
import("fmt")

func main(){

	//the compiler can infer the type of variables
	//it has a syntax for creating a variable and assigning to it in one line
	x := 34.0
	y := 43.0

	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("mean = %v, type of %T\n", mean, mean)
}

// $ go run mean.go
// x = 4, type of int
// y = 12, type of int
// mean = 8, type of int

// $ go run mean.go
// x = 4, type of float64
// y = 19, type of float64
// mean = 11.5, type of float64


