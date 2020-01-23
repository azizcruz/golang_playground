package main

import (
	"fmt"
	"strconv"
)

// define struct
type Person struct {
	// firstName string
	// lastName  string
	// gender    string
	// age       int

	firstName, lastName, gender string
	age                         int
}

// define a person method (value receiver)
func (p Person) greet() string {
	return "I'm " + p.firstName + " and I'm " + strconv.Itoa(p.age)
}

// define a person method (pointer receiver)
func (p *Person) increaseAge() {
	p.age++
}

// change name based on condition
func (p *Person) changeName(lastName string) {
	if len(p.firstName) > 3 {
		p.firstName = lastName
	}
}

func main() {
	// Init Person struct
	person1 := Person{
		firstName: "Aziz",
		lastName:  "Lioua",
		gender:    "m",
		age:       25}

	person1.increaseAge()
	person1.increaseAge()
	person1.increaseAge()

	person1.changeName("Ali")
	fmt.Println(person1)
	fmt.Println(person1.age)
	fmt.Println(person1.greet())
}
