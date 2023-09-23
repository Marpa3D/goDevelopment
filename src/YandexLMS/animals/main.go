// Интерфейсы. "Гошин питомец"
package main

import (
	"fmt"
	"os"
)

type Animal interface {
	MakeSound()
}

func MakeSound(a Animal) {
	a.MakeSound()
}

type Dog struct {
	sound string
}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

type Cat struct {
	sound string
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

func main() {
	cat1 := Cat{}
	cat1.MakeSound()
	dog1 := Dog{}
	dog1.MakeSound()

	fmt.Println("Для выхода из программы введите клавишу 'Q' и нажимте Enter")

	var exit rune

	fmt.Scanln(&exit)

	if exit == 'Q' || exit == 'q' {
		os.Exit(0)
	}
}
