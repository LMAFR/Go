package main

import "fmt"

func main() {

	var miCafe = Cafe{
		nombre: "expresso", // We could not specify the name of the variable in case you know the order in which these were written.
		precio: 3.20,
		azucar: false,
		leche:  0,
	}

	fmt.Print(miCafe)

}

type Cafe struct {
	nombre string
	precio float64
	azucar bool
	leche  int // "milk" would be the percentage of milk in your coffee.

}
