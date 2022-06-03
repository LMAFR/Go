package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Remember that a numeric variable whose value is not defined, has a default value of zero.
	var num int64
	rand.Seed(time.Now().Unix())
	// The first and third arguments of a for loop are optional in Go, so a while loop (which does not exist as statement in Go)
	// can be replicated as follows:
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}
}
