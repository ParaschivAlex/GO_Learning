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

/*func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}*/

func (englishBot) getGreeting() string { // nu pun (eb englishBot) pentru ca nu folosesc eb si pot lasa doar receiverul
	//costum logic here
	return "Hello man!"
}

func (spanishBot) getGreeting() string {
	//costum logic here
	return "Ola amigo!"
}
