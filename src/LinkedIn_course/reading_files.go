package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// We are going to use the library ioutil, in which we can define both data and possible errors as the output of the function to
	// read archives:
	data, error := ioutil.ReadFile("archivo.txt")

	// If we get some error, we will show it in terminal:
	print_error(error)

	// If we get no errors, we will show the data in terminal:
	fmt.Println(string(data))

}

func print_error(e error) { // "error" is also a type of variable

	if e != nil { // nil is similar to NULL in other languages, the condition means "if e has no value then..."
		panic(e) // panic is a function to stop the execution and print an error in terminal.
	}

}
