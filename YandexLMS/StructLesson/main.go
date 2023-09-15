// Работа с пакетами
package main

import (
	"fmt"
	students "github.com/Marpa3D/goDevelopment/YandexLMS/StructLesson/Students"
)

func main() {
	newStudent := students.Student{
		"Mark",
		45,
	}
	fmt.Println(newStudent)
}
