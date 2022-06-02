package main

import (
	"fmt"
	"strconv"
)

var integer16 int16 = 127
var integer32 int32 = 32767

func main() {
	// The "_" means that we are not going to use the value of that variable in this code. This is necessary because Go cannot compile programs which does not use all the defined variables.
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)

	fmt.Println(int32(integer16) + integer32)
	fmt.Println(i, s)
}
