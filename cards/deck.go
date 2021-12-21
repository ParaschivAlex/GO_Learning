package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings

type deck []string

func (d deck) print() { // d este o copie din deck, deck este tipul declarat mai sus
	for i, card := range d {
		fmt.Println(i, card)
	}
}
