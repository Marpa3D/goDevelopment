// Пакет Студенты
package Students

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) PrintDataStudent() {
	fmt.Printf("Имя студента: %s, Возраст: %d\n", s.Name, s.Age)
}
