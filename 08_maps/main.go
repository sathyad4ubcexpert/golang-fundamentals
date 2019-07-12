package main

import "fmt"

func main() {
	//Define Map
	emails := make(map[string]string)

	//Add key value pairs
	emails["ABC"] = "abc@email.com"
	emails["DEF"] = "def@email.com"
	emails["GHI"] = "ghi@email.com"
	emails["JKL"] = "jkl@email.com"

	fmt.Println(emails)

	// Define and add key value pairs
	contactNos := map[string]string{"sathya": "8688721439", "kalyani": "7989553520"}

	fmt.Println(contactNos)

	delete(emails, "JKL")

	fmt.Println(emails)
}
