package main

import "fmt"

func main() {
	x := 12
	y := 10

	//If else

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is not less than %d\n", x, y)
	}

	//Else if

	hairColor := "black"

	if hairColor == "gray" {
		fmt.Println("The hair color is GRAY")
	} else if hairColor == "white" {
		fmt.Println("The hair color is WHITE")
	} else {
		fmt.Println("The hair color is NEITHER GRAY NOR WHITE")
	}

	//Switch

	switch hairColor {
	case "gray":
		fmt.Println("The hair color is GRAY")
	case "white":
		fmt.Println("The hair color is WHITE")
	default:
		fmt.Println("The hair color is NEITHER GRAY NOR WHITE")
	}

}
