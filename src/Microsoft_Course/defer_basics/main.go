package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		// The deferred functions will be collected and run in opposite order after the regular statements of the loop had finished.
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
}
