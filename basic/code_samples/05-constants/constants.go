package main

import (
	"fmt"
)

// untyped constants
const a = 43
const b = 87.9003
const c = "Karthick there?"

// types constants
const d int = 43
const e float64 = 87.9003
const f string = "Karthick there?"

// other way to declar constants
const (
	g = 77
	h = 32.9090909
	i = "Yugin"
)

// iota

const (
	j int = iota
	k
	l
)

const (
	m = iota + 77
	n
	o
)

const (
	_ = iota // throw away the iota '0' value by ignoring it
	// kb = 1024
	kbi = 1 << (iota * 10) // iota==1, so shifts 10 binaries
	mbi = 1 << (iota * 10) // iota==2, so shifts 20 binaries
	gbi = 1 << (iota * 10) // iota==3, so shifts 30 binaries
)

func main() {
	// Constants check
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(g, h, i)

	// iota check
	fmt.Println(j, k, l)
	fmt.Println(m, n, o)

	// bit shifting
	m := 2
	fmt.Printf("Number %d's binary value is %b \n", m, m)

	n := m << 1
	fmt.Printf("Number %d's binary value is %b \n", n, n)

	// usage of just number
	fmt.Println("\n\nUsing numbers")
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Println("\nNumber \t\t\t\t Binary")
	fmt.Printf("%v \t\t\t\t %b\n", kb, kb)
	fmt.Printf("%v \t\t\t %b\n", mb, mb)
	fmt.Printf("%v \t\t\t %b\n", gb, gb)

	// using iota
	fmt.Println("\n\nUsing iota bit shifting <<")
	fmt.Println("\nNumber \t\t\t\t Binary")
	fmt.Printf("%v \t\t\t\t %b\n", kbi, kbi)
	fmt.Printf("%v \t\t\t %b\n", mbi, mbi)
	fmt.Printf("%v \t\t\t %b\n", gbi, gbi)

}
