package main

import "fmt"

func main() {

	x := 3
	y := 5

	if x > y {
		fmt.Printf("%d is greater than y\n", x)
	} else if y > x {
		fmt.Printf("%d is greater than x\n", y)
	} else {
		fmt.Printf("they are equal\n")
	}

	// Switch Case
	color := "Green"

	switch color {
	case "Blue":
		fmt.Println("Color is ", color)
	case "Red":
		fmt.Println("Color is ", color)
	default:
		fmt.Println("No color")
	}
}
