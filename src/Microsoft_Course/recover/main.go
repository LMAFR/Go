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
	defer func() {
		// Recover is always used in a defer function. It will return nil excepting the case in which a panic alert is raised.
		// If recover is in the code, the messages to debug the error, which caused the panic alert, will not be printed.
		// Recover and panic are a mechanism similar to try/catch in Python.
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
		// Observe the parenthesis added at the end of the next line. This seems to be related to the use of recover().
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}
