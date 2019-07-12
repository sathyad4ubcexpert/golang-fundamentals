package main

import "fmt"

func main() {
	a := 12345
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)
	fmt.Println(*&a)

	//change value with pointer
	*b = 10234516
	fmt.Println(a)

}
