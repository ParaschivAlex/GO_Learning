package main

import (
	"fmt"
	"net/http"
	"time"
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

	//fmt.Println(<-c) // primim o valoare de la channel
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " DOWN!")
		//c <- "Might be down!" //trimit mesaj catre server
		c <- link
		return
	}
	fmt.Println(link + " UP!")
	//c <- "Yea, it's up!"
	c <- link
	return
}
