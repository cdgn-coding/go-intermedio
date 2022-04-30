package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

// FullTimeEmployee implementa Printable
func (ft FullTimeEmployee) GetMessage() string {
	return fmt.Sprintf("FullTimeEmployee %v", ft)
}

// TemporaryEmployee implementa Printable
func (ft TemporaryEmployee) GetMessage() string {
	return fmt.Sprintf("TemporaryEmployee %v", ft)
}
