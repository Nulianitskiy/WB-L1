package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep

func mySleep(i time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*i)
	defer cancel()
	// Канал блокирует основную горутину, пока контекст не сработает
	<-ctx.Done()
	return

}

func main() {
	mySleep(5)
	fmt.Println("Cон закончился")
}
