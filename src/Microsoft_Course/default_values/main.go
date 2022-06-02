package main

import "fmt"

// Each type of integer has the same default value: zero.
var defaultInt int
var defaultfloat32 float32
var defaultFloat64 float64
var defaultBool bool
var defaultString string

func main() {
	fmt.Println(defaultBool, defaultFloat64, defaultInt, defaultString, defaultfloat32)
}
