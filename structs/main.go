package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Peterson",
		contactInfo: contactInfo{
			email:   "jim@disney.com",
			zipCode: 99035,
		},
	}

	jim.print()
	// jimPointer := &jim
	// jimPointer.updateName("Tony")
	jim.updateName("Frank")
	jim.print()
}

func (pointer *person) updateName(newFirstName string) {
	fmt.Println("1", pointer, &pointer, *pointer)
	(*pointer).firstName = newFirstName
	fmt.Println("2", pointer, &pointer, *pointer)
}

func (pointer *person) print() {
	// fmt.Printf("%+v", pointer)
	fmt.Println(&pointer)
	fmt.Println(*pointer)
}
