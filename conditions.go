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




}

// $ go run conditions.go
// idk how I feel about the absence of parentheses


