package main

import "fmt"

func main() {
	// Si no le mandas argument, es un canal sin buffer. Se bloquea hasta que alguien lea
	// Si se mandan mas mensajes que el tamaño del buffer, se bloquea hasta que alguien lo lea
	c := make(chan int, 1)
	c <- 1

	fmt.Println(<-c)
}
