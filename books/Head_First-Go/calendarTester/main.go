package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Best friend's birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019) // setter method for Date promoted to Event
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	// Getter methods for Date promoted to Event
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	// Can use dot operator chaining to call methods on the Date value directly
	fmt.Println(event.Date.Day())

	// Below code causes an error because Date struct fields are un-exported/protected
	// date2 := calendar.Date{year: 2021, month: 2, day: 10}
}
