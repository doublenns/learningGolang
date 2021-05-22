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
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 90210,
		},
	}

	fmt.Printf("%+v\n", jim)
}
