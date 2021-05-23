package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an English greeting
	return "Whaddup!"
}
func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a Spanish greeting
	return "Que paso!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
