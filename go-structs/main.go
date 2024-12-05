package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	jane := person{"Jane", "Doe"}
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Printf("%s %s\n", jane.firstName, jane.lastName)
	fmt.Printf("%s %s\n", alex.firstName, alex.lastName)
}
