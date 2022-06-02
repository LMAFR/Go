package main

import (
	"fmt"
	"os"
	"strconv"
)

// Be careful, this program requires two additional arguments to be run. For instance: "go run <path>/main.go 3 5"
func main() {
	// The arguments passed to the program using the command line are considered as strings
	sum, _ := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum: ", sum)
}

func calc(number1 string, number2 string) (sum int, mult int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mult = int1 * int2
	return
}
