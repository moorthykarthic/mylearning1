package main

import (
	"fmt"
)

func main() {

	// declare and assign a VALUE to String variable
	s := "Yugin Karthick"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// assign another value to the same variable
	s = "Nandhini Karthick"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// slice of bytes
	sliceOfBytes := []byte(s)
	fmt.Println(sliceOfBytes)
	fmt.Printf("%T\n", sliceOfBytes)

	//print as UTF-8 value
	for i := 0; i < len(sliceOfBytes); i++ {
		fmt.Printf("%#U\n", sliceOfBytes[i])
	}

	//print as hex char
	for j, val := range sliceOfBytes {
		fmt.Printf("\n position %d has the hex as %#x", j, val)
	}
	fmt.Println("")

	// numerical systems
	h := "H"
	fmt.Println(h)
	bs := []byte(h)
	fmt.Printf("%T\n", bs[0])  // type
	fmt.Printf("%b\n", bs[0])  // UTF-8 - bytes
	fmt.Printf("%#x\n", bs[0]) // hex char
}
