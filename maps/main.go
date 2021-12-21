package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(clrs map[string]string) {
	for key, value := range clrs {
		fmt.Println(key, "-->", value)
	}
}
