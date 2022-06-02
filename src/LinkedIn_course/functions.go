package main

import "fmt"

func main() {

	//fmt.Print(multiplicar(10, 8))
	//fmt.Print(ganado(1))

	say_hello("Alejandro", "Alvaro", "Guillermo", "Juan")

}

// Let's create a function which multiplies 2 values:
func multiplicar(a int, b int) int {
	// We could also have declared the argument types as "(a, b int)"
	// The type declared after parethesis is used to let GO know what type of data should return this function (the value of return, in the next
	// line)
	return a * b
}

// If we would like to return 2 values, we would write the function as follows:

func hola_mundo() (string, int) {
	var a string
	var b int

	a = "Hola mundo "
	b = 2

	return a, b
}

// Another way to write the last function would be the following one:

func hello_world() (a string, b float64) {

	a = "Hello world "
	b = 2.0

	// You have already declared the variables to return in the function definition, so you can simply write "return":
	return

}

// We could also define functions inside functions (closures):

func ganado(number int) (int, string) {

	// Let's create an anonymous function:
	vacas := func() int {
		return number * 10
	}

	// vacas is a function, so it will be called as vacas():
	return vacas(), " vacas"
}

// An interesting way to work in GO let us create an undefined number of arguments in a function using "...", let's try it!
func say_hello(friends ...string) {

	for _, name := range friends {
		fmt.Print("\nHello " + name) // \n is a lineskip.
	}

	// Pay attention to the omission of types for the return. We have not used return here, so we have not defined a type for it.
}
