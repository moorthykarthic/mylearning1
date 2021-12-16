package main

import "fmt"

// exercise 3
const h = "Karthick" //UNTYPED constant
const i int = 33     //TYPED constant

// exercise 6
const (
	o = 2021 + iota
	p = 2021 + iota
	q = 2021 + iota
	r = 2021 + iota
)

func main() {
	// exercise 1
	fmt.Println("\n============= Exercise - 1 =============")
	a := 55
	fmt.Printf("%v's decimal value is %d\n", a, a)
	fmt.Printf("%v's decimal value is %b\n", a, a)
	fmt.Printf("%v's decimal value is %#x\n", a, a)

	// exercise 2
	fmt.Println("\n============= Exercise - 2 =============")
	b := (10 == 15)
	c := (10 <= 15)
	d := (10 >= 15)
	e := (10 != 15)
	f := (10 < 15)
	g := (10 > 15)

	fmt.Printf("(10 == 15) \t--> %v\n", b)
	fmt.Printf("(10 <= 15) \t--> %v\n", c)
	fmt.Printf("(10 >= 15) \t--> %v\n", d)
	fmt.Printf("(10 != 15) \t--> %v\n", e)
	fmt.Printf("(10 < 15) \t--> %v\n", f)
	fmt.Printf("(10 > 15) \t--> %v\n", g)

	// exercise 3
	fmt.Println("\n============= Exercise - 3 =============")
	fmt.Printf("Value of UNTYPED constant --> %v\n", h)
	fmt.Printf("Value of TYPED constant --> %v\n", i)

	// exercise 4
	fmt.Println("\n============= Exercise - 4 =============")
	j := 10
	fmt.Printf("Decimal value of %v is\t : %d\n", j, j)
	fmt.Printf("Binary value of %v is\t : %b\n", j, j)
	fmt.Printf("Hex value of %v is\t : %#x\n\n", j, j)
	k := j << 2
	fmt.Printf("Decimal value of %v is\t : %d\n", k, k)
	fmt.Printf("Binary value of %v is\t : %b\n", k, k)
	fmt.Printf("Hex value of %v is\t : %#x\n\n", k, k)
	l := k >> 1
	fmt.Printf("Decimal value of %v is\t : %d\n", l, l)
	fmt.Printf("Binary value of %v is\t : %b\n", l, l)
	fmt.Printf("Hex value of %v is\t : %#x\n\n", l, l)
	m := l >> 1
	fmt.Printf("Decimal value of %v is\t : %d\n", m, m)
	fmt.Printf("Binary value of %v is\t : %b\n", m, m)
	fmt.Printf("Hex value of %v is\t : %#x\n", m, m)

	// exercise 5
	fmt.Println("\n============= Exercise - 5 =============")
	n := `Hi, I am 
	"Karthick 
	MOORTHY".
	It's another sentence`
	fmt.Println("Printing raw string literal : \n", n)

	// exercise 6
	fmt.Println("\n============= Exercise - 6 =============")
	fmt.Printf("%v\n%v\n%v\n%v\n", o, p, q, r)
}
