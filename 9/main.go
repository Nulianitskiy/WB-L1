package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5} // исходный массив чисел

	// Каналы для обмена данными
	dataCh := make(chan int)
	resultCh := make(chan int)

	// Запись чисел из массива в канал dataCh
	go func() {
		for _, x := range input {
			dataCh <- x
		}
		close(dataCh)
	}()

	// Умножение чисел из dataCh на 2 и запись результатов в канал resultCh
	go func() {
		for x := range dataCh {
			resultCh <- x * 2
		}
		close(resultCh)
	}()

	// Вывод результатов из канала resultCh в stdout
	for result := range resultCh {
		fmt.Println(result)
	}
}
