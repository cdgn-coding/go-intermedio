package main

import (
	"github.com/cdgn-coding/go-intermedio-concurrencia/8-modulos/utils"
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Carlos David")
	utils.Greet("Hello world from my module")
}
