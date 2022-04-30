package main

import (
	"fmt"
	"time"
)

func doSomething(i time.Duration, output chan<- int, param int) {
	time.Sleep(i)
	output <- param
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	duration1 := 4 * time.Second
	duration2 := 2 * time.Second

	go doSomething(duration1, channel1, 1)
	go doSomething(duration2, channel2, 2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-channel1:
			fmt.Println("Recibido de channel1", msg)
		case msg := <-channel2:
			fmt.Println("Recibido de channel2", msg)
		}
	}
}
