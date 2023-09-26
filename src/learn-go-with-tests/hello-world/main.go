// тест программы "Hello World"
package main

import "fmt"

const enHelloPrefix = "Hello, "

func Hello(s string) string {
	return enHelloPrefix + s + "!"
}

func main() {

	fmt.Println(Hello("Mark"))
}
