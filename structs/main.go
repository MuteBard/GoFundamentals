package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	//first way to initialize a struct
	// carl := {"Carl", "Severe"}
	// fmt.Println(carl)

	//second way to initialize a struct
	carl := person{firstName: "Carl", lastName: "Severe"}
	fmt.Println(carl)

	//third way to initialize a struct
	var erin person
	fmt.Println(erin)

	//updating a value
	carl.firstName = "carrrrl"

	//printing with the key : value
	fmt.Printf("%+v\n", carl)
	fmt.Printf("%+v\n", erin)

	metagross := person{
		firstName: "Meta",
		lastName:  "Gross",
		contact: contactInfo{
			email:   "metagross@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v\n", metagross)
}

// $ go run main.go
// {Carl Severe}
// { }
// {firstName:carrrrl lastName:Severe}
// {firstName: lastName:}
