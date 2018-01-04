package main

import "fmt"

func main() {
	fmt.Println("Full Cards\n===================")
	cards := newDeck()
	// cards.shuffle()
	cards.print()

	hands, ReminingCards := deal(cards, 5)
	fmt.Println("Card yang ada di tangan\n=======================")
	hands.print()
	fmt.Println("Ini untuk Remining Cards\n========================")
	ReminingCards.print()

	fmt.Println(hands.toString())
	fmt.Println(ReminingCards.toString())

	// Create File
	hands.createToFile("file/Hallo.txt")

	fmt.Println("Card di Read melalui file\n==========================")
	readFile := readFile("file/Hallo.txt")
	readFile.print()

}
