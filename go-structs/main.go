package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jane := person{"Jane", "Doe", contactInfo{"jane@test.com", 123345}}

	alex := person{firstName: "Alex", lastName: "Anderson", contact: contactInfo{email: "alex@test.com", zipCode: 243562}}

	john := person{}
	john.firstName = "John"
	john.lastName = "Doe"
	john.contact = contactInfo{email: "john@test.com", zipCode: 123456}

	fmt.Printf("%s %s, email: %s, zipCode: %d\n", jane.firstName, jane.lastName, jane.contact.email, jane.contact.zipCode)
	fmt.Printf("%s %s, email: %s, zipCode: %d\n", alex.firstName, alex.lastName, alex.contact.email, alex.contact.zipCode)
	fmt.Printf("%s %s, email: %s, zipCode: %d\n", john.firstName, john.lastName, john.contact.email, john.contact.zipCode)
}
