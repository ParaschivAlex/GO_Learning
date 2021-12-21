package main

import "fmt"

func main() {
	//var colors map[string]string
	colors := make(map[string]string)

	colors["new_color"] = "#23ff21"
	delete(colors, "new_color")

	fmt.Println(colors)
}
