package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	fmt.Println("Remaining in the Deck are:")
	remainingDeck.print()
}
