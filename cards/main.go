package main

func main() {
	// fmt.Println("Hello world!!")
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting))

	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("myCards")
	cards := newDeckFromFile("myCards")
	//cards.shuffleMyDeck()
	cards.shuffle()
	cards.print()
}
