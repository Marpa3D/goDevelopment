// ира "Карты"
package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard(), "Six of Spades"}

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
