package main
import("fmt")
func main(){
	for i := 0; i < 3; i++ {
		fmt.Printf("number is %v\n", i) 
	} 
}

// $ go run for.go
// number is 0
// number is 1
// number is 2

