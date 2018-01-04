package main

import "fmt"

var angka int // tarok di sini juga bisa variabel nya

//Data Type Golang
// -- bool = true, false
// -- string = "Hallo World"
// -- int = 0, -100000, 999999
// -- float64 = 10.0001, 0.00009, -10000.9

func main() {
	// Cara 1 bikin String
	// var card string = "Selama datang di cards"
	// Cara 2 bikin String
	// card := "Selamat Datang di Cards"

	card := "Selamat Datanag di Cards Session"
	angka = 10
	fmt.Println(angka)
	fmt.Println(card)
	fmt.Println(newCard())

	// ini cara call 2 function sekaligus, jadi command 'go run *.go or go run coba.go main.go'
	// printState()

	//Cara bikin Array []
	cardArray := []string{"Ace Of Diamonds", newCard()}
	cardArray = append(cardArray, "Queen Black Diamonds")
	fmt.Println()
	fmt.Println("Print Array CardArray")
	for i, cards := range cardArray { // ini cara for Array
		fmt.Println(i, cards)
	}
}

func newCard() string { // ini cara bikin Function
	king := "King Black Diamonds"
	return king
}
