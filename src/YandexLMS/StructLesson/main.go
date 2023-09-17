// Работа с пакетами
package main

import (
	"fmt"
	students "github.com/Marpa3D/goDevelopment/src/YandexLMS/StructLesson/Students"
	"github.com/Marpa3D/goDevelopment/src/YandexLMS/StructLesson/rectangle"
)

func main() {
	newStudent := students.Student{
		"Mark",
		45,
	}

	newStudent.PrintDataStudent()
	fmt.Printf("Тип студента: %T\n", newStudent)

	rec := rectangle.Rectangle{Heigth: 10, Width: 8}
	area := rec.Area()
	sW, sH := rec.Scale(8.0)
	fmt.Printf("%.2f, %.2f: %.2f\n", area, sW, sH)
}
