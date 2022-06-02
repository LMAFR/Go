package main

import "fmt"

func main() {

	// For loop in GO:

	for i := 1; i <= 10; i++ {
		fmt.Println("\nYour battery level is ", i*10, "%.") // Println creates a lineskip after each element printed.

		if i == 10 {
			fmt.Print("Charge completed!")
		}

	}

}
