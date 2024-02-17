package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 1, 2
	a, b = b, a

	c, d := 3, 4
	c, d = swap(c, d)

	fmt.Printf("a = %d, b = % d\n", a, b)
	fmt.Printf("c = %d, d = % d\n", c, d)
}
