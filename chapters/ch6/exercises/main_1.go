package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	person := makePerson("Cristhian", "Gómez", 99)
	personPtr := makePersonPointer("Cristhian", "Gómez", 1)
	fmt.Printf("main person = {%v}\n", person)
	fmt.Printf("main personPtr= {%v}\n", personPtr)
}

func makePerson(firstName string, lastName string, age int) Person {
	person := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	fmt.Printf("makePerson ={%v}\n", person)
	return person
}

func makePersonPointer(firstName string, lastName string, age int) *Person {
	personPtr := &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	fmt.Printf("makePersonPointer = {%v}\n", personPtr)
	return personPtr
}
