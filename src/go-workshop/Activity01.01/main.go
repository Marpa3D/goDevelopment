// Медицинская форма для кабинета врача. Работа с перменными
package main

import (
	"fmt"
	"time"
)

var (
	Debug     bool      = false
	Status    string    = "information"
	StartTime time.Time = time.Now()
)

func main() {

	name, lastName := "Jhon", "Calhoun"
	age := 48
	pinataAllergy := false

	fmt.Println(name, lastName, age, pinataAllergy)
	fmt.Println(Debug, Status, StartTime)

}
