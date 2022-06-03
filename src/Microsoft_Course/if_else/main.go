package main

import "fmt"

func somenumber() int {
	return -7
}

func main() {
	// The possibility of define variables into an if block is a property of the Go language.
	// Moreover, the "if" conditions in Go does not need brackets
	if num := somenumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// The next line of code would raise an error, because a variable defined in an if block is a local variable of that block.
	// fmt.Println(num)

}
