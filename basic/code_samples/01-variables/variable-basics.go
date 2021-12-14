package main

import (
	"fmt"
)

var k int     //declaring int at package scope
var l float32 // declaring float
var m bool    // declaring boolean
var s string  // nil for other types

//create your own TYPE
type kart int

var nan kart

func main() {
	fmt.Println("This is an another hello world")

	fmt.Printf("%T", k) //default value is 0
	fmt.Println(k)
	fmt.Printf("%T", l) //default value is 0.0
	fmt.Println(l)
	fmt.Printf("%T", m) //default value is false
	fmt.Println(m)
	fmt.Printf("%T", s) //default value is false
	fmt.Println(s)
	//assign raw string value with backtick key
	s = `Hi, I am "Karthick".`
	fmt.Println(s)

	//print my own type variable
	fmt.Printf("%T", nan)
	fmt.Println(nan, "nan value")

	// conversion -> convert your own type into the underlying type
	var yug int
	yug = int(nan)
	yug = yug + 1
	fmt.Printf("%T", yug)
	fmt.Println(yug, "Yug value")

	// traditional declaration
	var x int // try to use it for package level scope and for zero values
	x = 5
	fmt.Printf("%T", x)
	fmt.Println(x)

	// declare and assign value
	var y = 10
	fmt.Println(y)

	//short declaration
	z := 15
	fmt.Println(z)

	//Short declaration with expression
	a := 20 + 30
	fmt.Println(a)

	foo()

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Print(" : This is an even number")
		}
	}

	//catch error
	n, err := fmt.Println("Did you throw error", 55, false)
	fmt.Println(n)
	fmt.Println(err)

	//ignore error using '_' underscore
	h, _ := fmt.Println("Ignore the error returned", 55, false)
	fmt.Println(h)
	bar()

	fmt.Println("main is exited")
}

func foo() {
	fmt.Println("I am in foo function")
}

func bar() {
	fmt.Println("I am in BAR function")
}
