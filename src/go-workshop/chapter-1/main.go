// Код случайным образом выводит на консоль сообщение из заранее определенного списка
package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Привет",
	"Hello",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا‎",
	"Привет, мир",
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("За пределами слайса" + strconv.Itoa(index))
	}
	return helloList[index], nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(helloList))

	msg, err := hello(index)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
