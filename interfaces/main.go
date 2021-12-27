package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string { // nu pun (eb englishBot) pentru ca nu folosesc eb si pot lasa doar receiverul
	//costum logic here
	return "Hello man!"
}

func (sb spanishBot) getGreeting() string {
	//costum logic here
	return "Ola amigo!"
}
