package main

import "fmt"

//Удалить i-ый элемент из слайса.

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	originalSlice := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	modifiedSlice := removeIndex(originalSlice, indexToRemove)
	fmt.Println(modifiedSlice)
}
