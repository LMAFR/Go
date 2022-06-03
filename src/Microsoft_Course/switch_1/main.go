package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0, 1, 2:
		fmt.Println("Less than 3")
	case 3, 4, 5:
		fmt.Println("Less than 6 but more than 2.")
	default:
		fmt.Println("More than 5")
	}
}
