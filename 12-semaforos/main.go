package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Buffered channels como semaforos
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Println("Starting", i)
	time.Sleep(5 * time.Second)
	fmt.Println("Done", i)
	<-c
}
