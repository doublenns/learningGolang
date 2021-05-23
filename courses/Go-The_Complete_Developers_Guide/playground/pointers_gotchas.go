package main

import "fmt"

func main() {
	mySlice := []string{
		"Hi",
		"There",
		"How",
		"Are",
		"You",
	}

	updateSlice(mySlice)
	fmt.Println(mySlice)

	name := "doublenns"
	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
