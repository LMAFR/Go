package main

import "fmt"

func main() {

	// constants can be define as a variable, excepting ":=" notation. For example, I can define a constant like this:
	const constante int = 100
	// But I cannot define a constant like this: const constante := 100; or: constante := 100; (in the last case it would be interpreted as
	// a variable, nor a constant).

	fmt.Print(constante)
}
