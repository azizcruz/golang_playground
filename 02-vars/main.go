package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 and float64
	// complex64 and complex128

	// Using var
	var name = "Aziz"
	var age int = 25

	fmt.Println(name)
	fmt.Printf("%T %d\n", age, age)
}
