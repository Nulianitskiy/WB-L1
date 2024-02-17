package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBit(num int64, pos int, value int) int64 {
	mask := int64(1 << pos)
	if value == 1 {
		return num | mask
	} else {
		return num &^ mask
	}
}

func main() {
	var num int64 = 10
	pos := 0   // Номер бита (начиная с 0), который нужно установить
	value := 1 // Значение, которое нужно установить

	num = setBit(num, pos, value)
	fmt.Printf("Измененное число: %d\n", num)
}
