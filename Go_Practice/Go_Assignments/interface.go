package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main5() {
	rect := Rectangle{width: 10, height: 5}
	circ := Circle{radius: 7}

	shapes := []Shape{rect, circ}

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}
