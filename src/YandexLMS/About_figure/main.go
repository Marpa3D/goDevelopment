// Интрефейсы. "Все о фигуре"
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

func Area(a Shape) float64 {
	return a.Area()
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	fmt.Println("Все о фигуре")

	circle := Circle{radius: 45}
	area := circle.Area()

	rect := Rectangle{width: 14, height: 28}
	rArea := rect.Area()

	fmt.Printf("Circle area: %0.2f\nRectangle area: %0.2f", area, rArea)
}
