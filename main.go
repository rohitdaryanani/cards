package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of spades")
	// fmt.Println(cards)

	cards := newDeck()

	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
