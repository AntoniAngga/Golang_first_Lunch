package main

import (
	"fmt"
)

func main() {
	// Cara pertama bikin Object/map
	person := map[string]string{
		"Name":  "Antoni Angga",
		"Batch": "Happy Fox",
		"Class": "Full Stack Class",
	}
	fmt.Println(person)

	// Cara kedua bikin Object/map
	var person1 map[string]string
	fmt.Println(person1)

	// Cara ketiga bikin Object/map
	colors := make(map[int]string)

	colors[10] = "#fffff" // cara insert maps nya;

	delete(colors, 10) // cara delete, 2 parameter (variabel, dan key nya);

	printMap(person)
}

func printMap(c map[string]string) { // cara print maps
	for key, value := range c {
		fmt.Println(key, ":", value)
	}
}
