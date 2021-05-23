package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an English greeting
	return "Whaddup!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a Spanish greeting
	return "Que paso!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
