package main

import "fmt"

func main() {
	// Long way
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	println("=======")
	// Short way
	for i := 0; i < 10; i++ {
		println("Num ", i)
	}
	println("=======")
	// Fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			println("FizzBuzz ", i)
		} else if i%3 == 0 {
			println("Fizz ", i)
		} else if i%5 == 0 {
			println("Buzz ", i)
		} else {
			println(i)
		}
	}
}
