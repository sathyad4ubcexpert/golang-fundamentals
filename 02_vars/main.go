package main

import (
	"fmt"
)

var name = "Sathyanarayana Dudigama"
var age = 38
var weight float32 = 72.5

const isCool = true

func main() {

	//Shorthand var declaration
	height := 5.6
	email, contactno := "sathyad4ubcexpert@gmail.com", 8688721439

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Height: ", height)
	fmt.Println("Weight: ", weight)
	fmt.Println("Email: ", email, " Contact Number: ", contactno)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", height)
	fmt.Printf("%T\n", weight)

}
