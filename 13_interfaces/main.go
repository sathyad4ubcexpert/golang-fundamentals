package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {

	circle := Circle{x: 0, y: 0, radius: 12}
	rectangle := Rectangle{width: 6, height: 7}

	fmt.Printf("Circle Area is: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area is: %f\n", getArea(rectangle))

}
