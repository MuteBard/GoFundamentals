package main
import("fmt")
func main(){

	//dont have to type breaks anyomre, thank you, like why wasnt that always a thing tbh
	x := 2;
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("idc anymore")
	}


	y := "two";
	switch y {
	case "one":
		fmt.Println(1)
	case "two":
		fmt.Println(2)
	case "three":
		fmt.Println(3)
	case "four":
		fmt.Println(4)
	default:
		fmt.Println("69646320616e796d6f7265")
	}

}

// $ go run switch.go
// two
// 2