package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, number := range ints {
		if number%2 != 0 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}
	}
}
