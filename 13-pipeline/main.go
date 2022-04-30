package main

import "fmt"

func Generator(out chan<- int) {
	for i := 1; i < 10; i++ {
		out <- i
	}
	close(out)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print(in <-chan int) {
	for value := range in {
		fmt.Println(value)
	}
}

func main() {
	numbers := make(chan int, 2)
	doubles := make(chan int, 2)

	go Generator(numbers)
	go Double(numbers, doubles)
	Print(doubles)
}
