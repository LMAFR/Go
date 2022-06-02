package main

import "fmt"

var firstName string = "John"

func main() {
	lastName := "Doe"
	fullName := firstName + " " + lastName + " \t(alias \"Foo\")\n2021/04/23"
	fmt.Println(fullName)
}

// Note that the simple quotes (') are used for individual characters and runes, while double quotes (") are used for strings.
