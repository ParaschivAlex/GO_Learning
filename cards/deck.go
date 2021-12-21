package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
// newdeeck, print, shuffle, deal, savetofile, newdeckfromfile
type deck []string

func newDeck() deck { //cand se apeleaza functia, vom returna mereu ceva de tipul deck
	//nu adaugam receiver
	cards := deck{}

	cardSymbols := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four"}

	for _, symbol := range cardSymbols {
		for _, number := range cardNumbers {
			cards = append(cards, number+" of "+symbol)
		}
	}
	return cards
}

func (d deck) print() { // d este o copie din deck, deck este tipul declarat mai sus
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d))
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

func deal(d deck, handSize int) (deck, deck) { // deck,deck ii arata lui go ca vrem sa returnam 2 valori de tip deck
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //0666 oricine poate citi/scrie in fisier
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //bs = byteslice, err = error
	//daca totul merge ok, err == nil, altfel, ia o valoare de eroare
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	str := strings.Split(string(bs), ",") // transform in string sliceul de bytes, apoi transf stringul in slice
	return deck(str)
}
