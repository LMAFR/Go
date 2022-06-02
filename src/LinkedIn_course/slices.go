package main

import "fmt"

func main() {

	// A slice contains a dynamic amount of elements and is declared as an array without declare de number of elements it would contain:
	var fruits = []string{
		"Apples",
		"Grapes",
	}

	fmt.Print(fruits)

	// We can add elements using "append":
	fruits = append(fruits, "Pinepples")
	// And we can get the lenght of a slice using "len":
	fmt.Println(len(fruits), fruits[1:])
}
