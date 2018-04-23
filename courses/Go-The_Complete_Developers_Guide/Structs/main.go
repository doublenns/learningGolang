package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	claire := person{firstName: "Claire", lastName: "Luis"}
	fmt.Println(alex)
	fmt.Println(claire.firstName)

	var ben person
	fmt.Println(ben)
	fmt.Printf("%+v \n", ben)
	ben.firstName = "Ben"
	ben.lastName = "Phillips"
	fmt.Printf("%+v \n", ben)
}
