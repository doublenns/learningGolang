package main

func main() {
	cards := newDeck()

	//hand, remainingCards := deal(cards, 5)

	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards.shuffle()
	cards.print()
}
