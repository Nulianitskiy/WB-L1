package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	in := make(chan int)
	out := make(chan int)

	numbers := []int{2, 4, 6, 8, 10}

	go func() {
		for _, number := range numbers {
			in <- number
		}
		close(in)
	}()

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for n := range in {
				out <- n * n
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for n := range out {
		fmt.Println(n)
	}
}
