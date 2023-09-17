// Оператор type assertion
package main

import "fmt"

func printType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	fmt.Println("Opertator type assertion")

	printType("Это строка")
	printType(true)
	printType(88)
	printType(14.8)
	printType('c')

}
