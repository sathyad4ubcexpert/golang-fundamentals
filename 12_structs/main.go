package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// age       int
	// gender    string

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}
func main() {
	//Initialize person struct
	person1 := Person{firstName: "Sathya", lastName: "Dudigama", city: "Hyderabad", age: 38, gender: "Male"}

	//Alternative way to init person
	person2 := Person{"Dharmik", "Dudigama", "Hyd", "Male", 2}

	person3 := Person{"Kalyani", "Nagapuri", "Hyd", "Female", 35}

	fmt.Println(person1)
	fmt.Println(person2)
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person3.getMarried("Dudigama")
	fmt.Println(person3.greet())

}
