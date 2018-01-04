package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	numberCard := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}
	symbolCard := []string{"Hearts", "Diamond", "Spades", "Club"}

	for _, number := range numberCard {
		for _, symbol := range symbolCard {
			cards = append(cards, "The "+number+" of "+symbol)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
