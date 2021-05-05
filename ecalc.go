package main

import (
	"fmt"
	"math/rand"
	"math"
)

func Ecalc() {
	const num_of_tries int = 10000000

	var moyXnotries int

	for i := 0; i < num_of_tries; i++ {
		var (
			sum			float64
			numOfRand   int
		)
		for sum < 1 {
			sum += rand.Float64()
			numOfRand ++
		}
		moyXnotries += numOfRand
	}

	mye := float64( moyXnotries ) / float64(num_of_tries)

	error := math.Abs(mye - math.E)/math.E

	fmt.Println("My E  is", mye)
	fmt.Println("   E  is", math.E)
	fmt.Println("Error is", error)



}
