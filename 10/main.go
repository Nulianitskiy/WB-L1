package main

import (
	"fmt"
	"math"
)

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
//15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
//градусов. Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	// Группировка значений
	for _, temp := range temperatures {
		index := int(math.Floor(temp/10.0)) * 10
		groups[index] = append(groups[index], temp)
	}

	for key, value := range groups {
		fmt.Printf("%d: %v\n", key, value)
	}
}
