package main
import("fmt")
func main(){
	x := 10
	y := 20

	if x > 5 {
		fmt.Println("idk how I feel about the absence of parentheses")
	}else{
		fmt.Println("but ill get over it")
	}

	condition := x + y > 25 && y > 15 
	
	if condition {
		fmt.Println("its doing the thing")
		fmt.Printf("condition is type %T\n", condition)
	}

	//in the other direction, if can have an optional initialization statement
	if anothercondition := x + y; anothercondition == 30{
		fmt.Println("should be equal to 30")
	}

}

// $ go run conditions.go
// idk how I feel about the absence of parentheses

// $ go run conditions.go
// idk how I feel about the absence of parentheses
// its doing the thing
// condition is type bool

// $ go run conditions.go
// idk how I feel about the absence of parentheses
// its doing the thing
// condition is type bool
// should be equal to 30

