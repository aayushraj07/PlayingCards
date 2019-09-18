package main

func main() {
	cards := deck{"Ace of Diamonds", "Five of Spades", newCards()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCards() string {
	return "Five of Diamonds"
}
