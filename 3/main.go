package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
//квадратов(2*2+3*3...) с использованием конкурентных вычислений.

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

	s := 0
	for n := range out {
		s += n
	}
	fmt.Println(s)
}
