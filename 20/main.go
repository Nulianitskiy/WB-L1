package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func reverseWords(str string) string {
	words := strings.Fields(str)
	reversed := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}
	return strings.Join(reversed, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}
