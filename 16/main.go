package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivotIndex := len(arr) / 2
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	quicksort(arr[:left])
	quicksort(arr[left+1:])
}

func main() {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	fmt.Println("Исходный массив:", arr)
	quicksort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
