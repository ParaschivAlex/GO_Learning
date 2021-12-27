package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //declarare channel

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c) // primim o valoare de la channel
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " DOWN!")
		c <- "Might be down!" //trimit mesaj catre server
		return
	}
	fmt.Println(link + " UP!")
	c <- "Yea, it's up!"
	return
}
