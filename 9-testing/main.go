package main

import (
	"fmt"
	"math"
)

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func ComputeOnceFibonacci(n int) int {
	size := math.Max(2, float64(n+1))
	values := make([]int, int(size))
	values[0] = 0
	values[1] = 1

	for i := 2; i <= n; i++ {
		values[i] = values[i-1] + values[i-2]
	}

	return values[n]
}

func main() {
	fmt.Println("2 + 2 =", Sum(2, 2))

	// go test -cpuprofile=cpu.out
	// go tool pprof cpu.out
}
