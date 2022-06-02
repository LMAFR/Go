package main

import (
	"fmt"
)

func main() {

	// First way to declare a variable:
	var name string
	// And assign a value to it:
	name = "mundo"

	// Second way to do both things:
	var ending string = "!"

	// Third way to do it:
	greet := "Hola " + name + ending

	fmt.Print(greet)
}
