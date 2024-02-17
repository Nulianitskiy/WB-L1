package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func checkUnique(str string) bool {
	seen := make(map[string]bool)
	letters := strings.Split(strings.ToLower(str), "")

	for _, l := range letters {
		if _, ok := seen[l]; ok {
			return false
		}
		seen[l] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("\"%s\" уникальность символов: %t\n", str1, checkUnique(str1))
	fmt.Printf("\"%s\" уникальность символов: %t\n", str2, checkUnique(str2))
	fmt.Printf("\"%s\" уникальность символов: %t\n", str3, checkUnique(str3))
}
