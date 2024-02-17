package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры
//Human (аналог наследования).

type Human struct {
	name string
}

func (h *Human) SaySmth(text string) {
	fmt.Printf("%s: %s\n", h.name, text)
}

type Action struct {
	Human
	action string
}

func (a *Action) DoSmth() {
	fmt.Printf("%s %s\n", a.name, a.action)
}

func main() {
	detective := Action{
		Human: Human{
			name: "Harry",
		},
		action: "arriving on the scene"}

	detective.DoSmth()
	detective.SaySmth("Sunrise, Parabellum")
}
