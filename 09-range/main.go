package main

import "fmt"

func main() {
	ids := []int{33, 45, 65, 78, 67, 63, 42, 43}

	// Loop through ids using range
	for i, id := range ids {
		fmt.Println(i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("id: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum ", sum)

	// Range with map
	names := map[string]int{
		"aziz": 25,
		"ali":  28}

	for k, v := range names {
		fmt.Printf("%s => %d\n", k, v)
	}
}
