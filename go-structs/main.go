package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	jane := person{"Jane", "Doe"}

	alex := person{firstName: "Alex", lastName: "Anderson"}

	john := person{}
	john.firstName = "John"
	john.lastName = "Doe"

	fmt.Printf("%s %s\n", jane.firstName, jane.lastName)
	fmt.Printf("%s %s\n", alex.firstName, alex.lastName)
	fmt.Printf("%s %s\n", john.firstName, john.lastName)
}
