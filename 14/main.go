package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}

func main() {
	checkType(5)
	checkType("hello")
	checkType(true)

	ch := make(chan int)
	checkType(ch)
}

func checkType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Printf("Переменная %v имеет тип int\n", value)
	case string:
		fmt.Printf("Переменная %v имеет тип string\n", value)
	case bool:
		fmt.Printf("Переменная %v имеет тип bool\n", value)
	case chan int:
		fmt.Printf("Переменная %v имеет тип channel\n", value)
	default:
		fmt.Printf("Неизвестный тип переменной: %v\n", reflect.TypeOf(value))
	}
}
