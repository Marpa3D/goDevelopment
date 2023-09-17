// Структуры данных. Массив
package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var someArray [5]int

	fmt.Printf("Адрес массива: %p\nЗначения в массиве: %+v\nТип %T\n", &someArray, someArray, someArray)
}
