package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverseString(s string) string {
	letters := strings.Split(s, "")
	reversed := make([]string, 0, len(letters))
	for i := len(letters) - 1; i >= 0; i-- {
		reversed = append(reversed, letters[i])
	}
	return strings.Join(reversed, "")
}

func main() {
	str := "главрыба"
	fmt.Println(reverseString(str))
}
