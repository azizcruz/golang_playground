package main

import "fmt"

func main() {
	// Arrays
	var fruits [2]string
	// Assign
	fruits[0] = "Apple"
	fruits[1] = "Grapes"

	fmt.Println(fruits)

	// Declare and Assign
	names := [2]string{"Aziz", "Ali"}

	fmt.Println(names)

	// Slices
	cars := []string{"BMW", "TOYOTA", "MERCEDES"}

	fmt.Println(cars)

}
