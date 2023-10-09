package main

func main() {
	//cards := newDeckFromFile("mydeck.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards.saveToFile("mydeck.txt")
}
