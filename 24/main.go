package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками, которые
//представлены в виде структуры Point с инкапсулированными параметрами x,y и
//конструктором.

type Point struct {
	x, y float64
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	// Создание двух точек
	point1 := Point{3, 4}
	point2 := Point{0, 0}

	// Вычисление расстояния между точками
	distance := Distance(point1, point2)
	fmt.Printf("Расстояние между точками: %f\n", distance)
}
