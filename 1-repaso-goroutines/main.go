package main

import (
	"fmt"
	"time"
)

func doSomething(c chan<- int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 0
}

func main() {
	c := make(chan int)
	go doSomething(c)
	// Unpack
	<-c
}
