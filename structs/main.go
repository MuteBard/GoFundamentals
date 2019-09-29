package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	metagross := person{
		firstName: "Meta",
		lastName:  "Gross",
		contactInfo: contactInfo{
			email:   "metagross@gmail.com",
			zipCode: 94000,
		},
	}

	metagross.updateName("MegaMeta")
	metagross.println()
}

//As the code is set up now, this will not work as intended, we need to talk about pointers
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) println() {
	fmt.Printf("%+v\n", p)
}
