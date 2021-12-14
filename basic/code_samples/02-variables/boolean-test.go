package main

import "fmt"

var x bool

func main() {

	fmt.Println(x)
	x = true
	fmt.Println(x)

	y := 4
	z := 17
	fmt.Println(y, " > ", z, ":", y > z)
	fmt.Println(y, " == ", z, ":", y == z)
	fmt.Println(y, " < ", z, ":", y < z)
	fmt.Println(y, " >= ", z, ":", y >= z)
	fmt.Println(y, " <= ", z, ":", y <= z)
	fmt.Println(y, " != ", z, ":", y != z)

	// The imitation movie
}
