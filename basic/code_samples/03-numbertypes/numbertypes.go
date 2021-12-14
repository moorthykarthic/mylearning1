package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

func main() {
	x = 77
	y = 8.480001

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// unsigned integer
	var a uint8 = 255    // 0 to 255
	var b uint16 = 65535 // 0 to 65535
	var c uint32 = 127   // 0 to 2 power 32
	var d uint64 = 127   // 0 to 2 power 64

	fmt.Println(a, b, c, d)

	// signed integer
	var e int8 = -128    // -128 to 127
	var f int16 = -32768 // -32768 to 32767
	var g int32 = 127    // -2 power 32 to 2 power 32
	var h int64 = 127    // -2 power 64 to 2 power 64
	fmt.Println(e, f, g, h)

	//byte

	//rune

	//uint // either 32 or 64 bits
	//int //same size as unit

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
