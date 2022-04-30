package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	// Funciones de un solo uso
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	// Funciones anonimas para go routines
	c := make(chan int)
	go func() {
		fmt.Println("Simulando proceso largo")
		time.Sleep(5 * time.Second)
		fmt.Println("Terminando proceso largo")
		c <- 0
	}()
	<-c
}
