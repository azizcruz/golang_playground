package main

import "fmt"

func sayHello(name string) string {
	return "Hello " + name
}

func add(n1 int, n2 int) int {
	return n1 + n1
}

func main() {
	fmt.Println(sayHello("aziz"))
	fmt.Println(add(3, 5))
}
