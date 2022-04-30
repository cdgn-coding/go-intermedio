package main

import "fmt"

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(workerId int, in <-chan int, out chan<- int) {
	var result int
	for n := range in {
		fmt.Println("Worker", workerId, "Starting to compute", n)
		result = Fibonacci(n)
		fmt.Println("Worker", workerId, "Computed Fibonacci of", n, "to be", result)
		out <- result
	}
}

func main() {
	tasks := []int{2, 3, 4, 5, 10, 20, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))

	results := make(chan int, len(tasks))

	for i := 0; i <= nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for i := 0; i < len(tasks); i++ {
		<-results
	}
	close(results)
}
