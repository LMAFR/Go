package main

import "fmt"

func main() {
	switch num := 15; {
	case num < 50:
		// Observe we use Printf to be able to replace %d by a numeric varible
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is more than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	}
}

// Observe that the result also return the print for the >100 condition, but it is not met. The key point here is that the fallthrough statement go down to
// the next case statement and go through its lines without check if the case condition is met. Because of that, we have to be careful with the fallthrough
// statement.
