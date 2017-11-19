package main

func main() {
	cards := newDeck()
	hand, remaingingDeck := deal(cards, 5)
	hand.print()
	remaingingDeck.print()
}
