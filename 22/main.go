package main

import (
	"fmt"
	"math"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две
//числовых переменных a,b, значение которых > 2^20.

func main() {
	a := math.Pow(2, 21)
	b := math.Pow(2, 22)

	multiplication := a * b
	division := b / a
	addition := a + b
	subtraction := b - a

	fmt.Printf("a: %f\n", a)
	fmt.Printf("b: %f\n", b)
	fmt.Printf("Умножение: %f\n", multiplication)
	fmt.Printf("Деление: %f\n", division)
	fmt.Printf("Сложение: %f\n", addition)
	fmt.Printf("Вычитание: %f\n", subtraction)
}
