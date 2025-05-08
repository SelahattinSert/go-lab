package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()

	cards.shuffle()

	fmt.Println("Shuffled deck: ")
	cards.print()

	filename := "my_cards.txt"
	err := cards.saveToFile(filename)
	if err != nil {
		panic("Exception occurred while saving file.")
	}

	loadedDeck := newDeckFromFile(filename)

	fmt.Println("Loaded deck: ")
	loadedDeck.print()

	hand, remaining := deal(loadedDeck, 4)

	fmt.Println("Hand: ")
	hand.print()

	fmt.Println("Remaining deck: ")
	remaining.print()
}
