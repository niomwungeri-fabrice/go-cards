package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("cards")
	cards := loadFileFromFile("cards")
	// cards.print()
	// fmt.Println("----------------")
	cards.shuffle()
	cards.print()
}
