package main

import "fmt"

func main() {
	a := 5
	b := &a

	// b is the address in the memory where a is stored
	fmt.Println(a, b)

	// check the type of a and b
	fmt.Printf("%T %T \n", a, b)

	// if we want to read what is inside this address we use *
	fmt.Println(*b)

	// change value by reference
	*b = 10
	fmt.Println(a)

}
