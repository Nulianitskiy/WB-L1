package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func intersection(set1, set2 []int) []int {
	set := make(map[int]bool)
	var intersect []int

	for _, val := range set1 {
		set[val] = true
	}

	for _, val := range set2 {
		if set[val] {
			intersect = append(intersect, val)
		}
	}

	return intersect
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersection(set1, set2)
	fmt.Println("Пересечение:", result)
}
