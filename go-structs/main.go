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

	//	alex := person{firstName: "Alex", lastName: "Anderson", contact: contactInfo{email: "alex@test.com", zipCode: 243562}}

	john := person{}
	john.firstName = "John"
	john.lastName = "Doe"
	john.contact = contactInfo{email: "john@test.com", zipCode: 123456}

	jane.print()

	jane.updateName("Susan")
	jane.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%s %s, email: %s, zipCode: %d\n", p.firstName, p.lastName, p.contact.email, p.contact.zipCode)
}
