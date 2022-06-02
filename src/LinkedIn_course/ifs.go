package main

import (
	"fmt"
)

func main() {

	// If condition in GO:
	puntos := 10

	if puntos == 10 {
		fmt.Print("You got the best score!")
	} else if puntos >= 5 {
		fmt.Print("You still could improve, keep going!")
	} else {
		fmt.Print("Study more for the next exam!")
	}

	// Other logic operators in GO: >, <, >=, <=, !=,

}
