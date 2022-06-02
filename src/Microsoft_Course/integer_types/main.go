package main

import "fmt"

var integer8 int8 = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807

func main() {
	fmt.Println(integer8, integer16, integer32, integer64)
	rune := 'G'
	fmt.Println(rune)
	// Error produced due to operate using two different types of integers:
	// fmt.Println(integer8 + integer16)
}
