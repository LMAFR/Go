package main

import "fmt"

// threshold for int32 type:
var integer32_limit int32 = 2147483647
var integer32_overflow int32 = 2147483648

func main() {
	fmt.Println(integer32, integer32_overflow)
}
