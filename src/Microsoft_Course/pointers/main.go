package main

import "fmt"

func main() {
	firstName := "John"
	// The parameter "name" of the updateName function is a local variable for that function, so it will not replace the previous value of firstName in usual conditions.
	// However, if we add & to the left of the variable that we want to update, it will be considered as a memory pointer, not a global variable
	updateName(&firstName)
	fmt.Println(firstName)
}

func updateName(name *string) {
	// If we want to update the value for the variable used as the parameter "name", we will have to use a pointer in that variable (firstName in this case), but also
	// we have to show that name will be the parameter whose value will replace the pointed variable. We do that using asterisk to the left of the type of the parameter and
	// to the left of the parameter when it is updated with a value.
	*name = "David"
}
