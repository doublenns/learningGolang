package main

import (
	"fmt"
)

func newCard() string {

	return "Ace of Spades"

}

func main() {
	// var card string = "Ace of Spades"

	card := newCard()
	fmt.Println(card)
}
