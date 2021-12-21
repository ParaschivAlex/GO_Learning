package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	/*alex := person{"Alex", "Anderson"}
	bob := person{firstName: "Bob", lastName: "Bobderson"}
	fmt.Println(alex)
	fmt.Println(bob)*/

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Bob"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
