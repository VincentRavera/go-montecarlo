package main

import (
	"fmt"
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
func main() {
	var sep string = "==========================================="
	fmt.Println(sep)
	fmt.Println("Caluclting pi")
	Picalc()
	fmt.Println(sep)
	fmt.Println("Caluclting e")
	Ecalc()
	fmt.Println(sep)
}
