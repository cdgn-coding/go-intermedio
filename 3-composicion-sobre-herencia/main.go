package main

import "fmt"

func main() {
	// El concepto de herencia no existe
	// Se prefiere la composicion sobre la herencia
	ft := FullTimeEmployee{}
	ft.name = "Carlos David"
	ft.age = 27
	ft.id = 1
	fmt.Println(ft)
	GetMessage(ft.Person)
}
