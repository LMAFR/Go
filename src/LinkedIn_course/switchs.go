package main

import (
	"fmt"
	"time"
)

func main() {

	// The "switch" controller:

	current_date := time.Now()

	switch current_day := current_date.Weekday(); current_day {

	case 0:
		fmt.Println(("\nToday is Sunday!"))
	case 6:
		fmt.Println("\nToday is Saturday!")
	default:
		fmt.Println("\nHOLY S... I am 5 minutes late to work!")

	}

}
