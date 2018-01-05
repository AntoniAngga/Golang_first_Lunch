package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func main() {
	var angga = person{
		firstName: "Antoni",
		lastName:  "Angga",
		age:       24,
		contactInfo: contactInfo{
			email:   "antoniangga14@gmail.com",
			zipcode: 1090,
		},
	}

	angga.print()

	anggaPointer := &angga          // bikin variable yang isi dengan struct itu
	anggaPointer.updateName("Coba") // dengan variable update name

	fmt.Printf("\n")
	angga.print()

	temp := "asdasdasd"
	fmt.Println(*&temp)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
