// Передача интерфейсов
package main

import "fmt"

type Shape interface {
	Area() float64
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

type Rect struct {
	Width  float64
	Heigth float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Heigth
}

func main() {
	rect := Rect{10.0, 14.8}
	area := calculateArea(rect)
	fmt.Println(area)
}
