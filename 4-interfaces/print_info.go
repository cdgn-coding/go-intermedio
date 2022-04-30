package main

import "fmt"

type Printable interface {
	GetMessage() string
}

func PrintInfo(p Printable) {
	fmt.Println(p.GetMessage())
}
