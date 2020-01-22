package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign value keys
	emails["aziz"] = "aziz@gmail.com"
	emails["ali"] = "ali@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "ali")
	fmt.Println(emails)

	// Declare map and add key values
	names := map[string]int{
		"aziz": 25,
		"ali":  28}

	fmt.Println(names)
}
