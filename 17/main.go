package main

import (
	"fmt"
)

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("%d найден в массиве arr на позиции %d\n", target, index)
	} else {
		fmt.Printf("%d не найден в массиве arr\n", target)
	}
}
