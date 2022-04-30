package main

import "fmt"

// Funcion variadica con operador ...
func sum(values ...int) int {
	var total int
	for _, value := range values {
		total += value
	}
	return total
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	// LLamando funcion variadica o con numero variable de argumentos
	fmt.Println(sum(2, 2))
	fmt.Println(sum(2, 2, 10, 40))

	// Named return
	fmt.Println(getValues(2))
}
