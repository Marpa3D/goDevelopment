// Бинарный поиск
package main

import (
	"fmt"
	"sort"
	"strconv"
)

// binarySearch - функция бинарного поиска в массиве из целых чисел
func binarySearch(arr []int, target int) (int, bool) {

	sort.Ints(arr) // первый шаг - сортировка массива

	var mid int // середина и текущее число

	min := 0
	high := len(arr) - 1

	for min <= high {
		mid = (min + high) / 2

		if mid == target {
			return mid, true
		}
		if mid > target {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return 0, false
}

func main() {

	var numElem int

	fmt.Println("Введите колличество эллементов массива:")
	fmt.Scan(&numElem)

	array := make([]int, numElem)
	fmt.Println("Создается массив...")
	for i := 0; i < numElem; i++ {
		array[i] = i
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()

	var lookNum int
	fmt.Println("Какой элемент будем искать?")
	fmt.Scan(&lookNum)

	nowNum, res := binarySearch(array, lookNum)
	if res {
		fmt.Printf("Найден элемент: %v", strconv.Itoa(nowNum))
	} else {
		fmt.Println("Ничего не найдено.")
	}

}
