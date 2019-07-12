package main

import "fmt"

func main() {
	ids := []int{12, 45, 76, 43, 87, 9, 765}

	// Loops through range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum:", sum)

	// Loops through map
	emails := map[string]string{"sathya": "sathyad4ubcexpert@email.com", "ABC": "abc@email.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}
