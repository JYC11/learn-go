package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jimson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}
	jim.updateFirstName("jimmy")
	jim.print()
}

func (p *person) updateFirstName(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
