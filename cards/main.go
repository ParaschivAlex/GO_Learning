package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	//fmt.Println(cards.toString())
	//hand, remainingDeck := deal(cards, 5)
	//	hand.print()
	//remainingDeck.print()
}
