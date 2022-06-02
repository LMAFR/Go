package main

import "fmt"

func main() {

	// First way to define an array in GO:
	var array1 [2]string
	array1[0] = "Apples"
	array1[1] = "Grapes"

	// Second way to define an array in GO:
	array2 := [2]string{
		"Football",
		"Basketball",
	}

	// We could build an array composed by arrays:
	var fruits [2][2]string

	// First fruit
	fruits[0][0] = "Golden"
	fruits[0][1] = "Apple"

	// Second fruit
	fruits[0][0] = "Green"
	fruits[0][1] = "Grape"

	fmt.Print(array1, array2, fruits[0], fruits[0][0])

}
