package main

import "fmt"

// As the uint type cannot have sign, the next definition will raise an error
var neg_uint uint = -10

func main() {
	fmt.Println(neg_uint)
}
