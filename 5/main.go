package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться.

func main() {
	ch := make(chan int)
	var N time.Duration = 3

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	go func() {
		for {
			select {
			case n, ok := <-ch:
				if !ok {
					fmt.Println("Канал закрыт")
					return
				}
				fmt.Println("Принято значение:", n)
			case <-time.After(3 * time.Second):
				fmt.Println("Время ожидания данных истекло")
				return
			}
		}
	}()

	time.Sleep(N * time.Second) // ждем 6 секунд, чтобы программа успела завершиться
	fmt.Println("Время истекло")
}
