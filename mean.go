package main
import("fmt")

func main(){
	var x float64
	var y float64

	x = 4
	y = 19

	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	var mean float64
	mean = (x + y) / 2
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


