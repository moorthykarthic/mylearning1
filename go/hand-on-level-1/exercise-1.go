package main

import (
	"fmt"
)

// exercise 2
var a int
var b string
var c bool

// exercise 3
var d int = 42
var e string = "James Bond"
var f bool = true

//exercise 4
type kart int

var g kart

//exercise 5
var h int

func main() {

	// exercise 1
	fmt.Println("=======Exercise : 1")
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)

	// exercise 2
	fmt.Println("=======Exercise : 2")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// exercise 3
	fmt.Println("=======Exercise : 3")
	s := fmt.Sprintf("%v\t%v\t%v\t", d, e, f)
	fmt.Println(s)

	//exercise 4
	fmt.Println("=======Exercise : 4")
	fmt.Println(g)
	fmt.Printf("%T\n", g)
	g = 73000
	fmt.Println(g)

	//exercise 5
	fmt.Println("=======Exercise : 5")
	h = int(g)
	fmt.Println(h)
	fmt.Printf("%T\n", h)

}
