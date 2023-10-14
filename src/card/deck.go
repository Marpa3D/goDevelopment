package main

import "fmt"

// Создание типа "Колода" (Deck). Это слайс типа string
type deck []string

func NewDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Пик", "Бубен", "Червей", "Треф"}
	cardValues := []string{"Туз", "Двойка", "Тройка", "Четверка"}

	for _, suits := range cardsSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" "+suits)
		}
	}
	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
