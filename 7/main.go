package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

type SafeMap struct {
	mu   sync.Mutex
	data map[string]string
}

func (sm *SafeMap) Set(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (string, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, ok := sm.data[key]
	return value, ok
}

func main() {
	sm := SafeMap{data: make(map[string]string)}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		sm.Set("key1", "value1")
	}()

	go func() {
		defer wg.Done()
		value, ok := sm.Get("key1")
		if ok {
			fmt.Println("Got value:", value)
		}
	}()

	wg.Wait()
}
