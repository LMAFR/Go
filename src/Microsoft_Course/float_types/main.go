package main

import (
	"fmt"
	"math"
)

var Float32 float32 = 21474
var Float64 float64 = 922337

const e = 2.71828
const Avogadro = 6.02214129e23
const Planck = 6.62606957e-34

func main() {
	fmt.Println(Float32, Float64)
	// The threshold for these types are
	fmt.Println(math.MaxFloat32, math.MaxFloat64)
	// The type of the values whose type is not specified are deduced by Go
	fmt.Println(e, Avogadro, Planck)
}
