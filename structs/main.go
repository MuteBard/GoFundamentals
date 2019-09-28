package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	// one way to declare a struct
	carl := person{firstName: "Carl", lastName: "Severe"}
	fmt.Println(carl)

	// second way to declare a struct
	var erin person
	fmt.Println(erin)

	//updating a value
	carl.firstName = "carrrrl"

	//printing with the key : value
	fmt.Printf("%+v\n", carl)
	fmt.Printf("%+v\n", erin)
}

// $ go run main.go
// {Carl Severe}
// { }
// {firstName:carrrrl lastName:Severe}
// {firstName: lastName:}
