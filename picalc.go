package main

import (
	"fmt"
	"math"
	"math/rand"
)

//
// ACercl = pi * r*r
// ACarr  =  (2 * r )**2 = 4 r*r
//
// ACercl = pi * (r**2 * 4 ) / 4
// ACercl = pi * ACarr / 4
//
// pi = 4 * ACercl / ACarr
//
//
func Picalc() {
	const num_of_point int = 10000000
	var num_in_circle int = 0

	for i := 0; i < num_of_point; i++ {
		var x = rand.Float64()
		var y = rand.Float64()

		if x*x+y*y < 1 {
			num_in_circle += 1
		}

	}
	mypi := 4.0 * float64(num_in_circle) / float64(num_of_point)

	error := float64(math.Abs(mypi-math.Pi ) / math.Pi)
	fmt.Println("My PI is", mypi)
	fmt.Println("   PI is", math.Pi)
	fmt.Println("Error is", error)

}
