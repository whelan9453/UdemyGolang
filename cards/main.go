package main

import (
	"fmt"
)

func main() {
	cards := newDeckFromFile("myCards")
	fmt.Println(cards)
	cards.shuffle()
	fmt.Println(cards)
}
