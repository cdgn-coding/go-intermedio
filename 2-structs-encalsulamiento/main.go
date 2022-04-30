package main

import "fmt"

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)
	e.id = 1
	e.name = "Name"
	fmt.Printf("%v\n", e)
	// Receiver functions
	e.SetId(5)
	e.SetName("Other Name")
	fmt.Println(e)
	fmt.Println(e.GetName())
	fmt.Println(e.id)

	// Construccion de structs especifcando campos
	e2 := Employee{
		id:   1,
		name: "Ramon",
	}

	fmt.Println(e2)

	// Usando New pone todo en zero values
	e3 := new(Employee)
	fmt.Println(*e3)

	// Creando una funcion constructora
	e4 := NewEmployee(1, "Ramon")
	fmt.Println(*e4)
}
