// Интерфейсы
package main

import "fmt"

type Human interface {
	SayHi()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHi() {
	fmt.Printf("Человек по имени %s, а возраст его: %d.\n", p.Name, p.Age)
}

func main() {
	var h Human = Person{
		Name: "Mark",
		Age:  45,
	}
	h.SayHi()
	fmt.Println("End")
}
