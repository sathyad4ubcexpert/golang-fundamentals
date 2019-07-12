package main

import "fmt"

func greeting(name string) string {
	return "Hello, " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {

	fmt.Println(greeting("Sathyanarayana"))
	fmt.Println(getSum(6543, 7654))

}
