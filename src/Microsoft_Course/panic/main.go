package main

import "fmt"

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		// Abort the execution with a panic alert just before the deferred lines accumulated.
		// It also adds some messages at the end of the execution in order to make it easier to debug the error.
		panic("highlow() low greater than or equal to high.")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	// This recursive statement will ensure we raise the panic alert in some moment
	highlow(high, low+1)

}

func main() {
	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}
