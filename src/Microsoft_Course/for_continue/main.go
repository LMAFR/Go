package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		if i%5 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is ", sum)
}
