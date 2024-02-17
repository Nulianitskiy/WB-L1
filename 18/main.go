package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика

type Counter struct {
	mu      sync.Mutex
	counter int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup
	numIncrements := 100

	for i := 0; i < numIncrements; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
