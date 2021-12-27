package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := ioutil.ReadFile(os.Args[1])
	fmt.Println(string(f))
}
