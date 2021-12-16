package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("\n============= for loop =============")
	// for init; condition; post
	for i := 0; i < 5; i++ {
		fmt.Printf("Outer loop : %v\n", i)
		// Nested loop
		for j := 0; j < 2; j++ {
			fmt.Printf("\tInner loop : %v\n", j)
		}
	}

	// for statement
	x := 1
	fmt.Println("\n============= for statement =============")
	for x <= 5 {
		fmt.Printf("Printing x value : %v\n", x)
		x = x + 1
	}

	// for statement with if condition
	y := 0
	fmt.Println("\n============= for statement with if condition =============")
	for {
		y = y + 1
		if y > 50 {
			break
		}
		if y%5 != 0 {
			continue
		}
		// below statemets will be continues
		fmt.Printf("Printing y value : %v\n", y)
		// for j := 0; j < 2; j++ {
		// 	fmt.Printf("\tInner loop : %v\n", j)
		// }

	}

	// Conditional - if statement
	fmt.Println("\n============= if statement =============")
	if true {
		fmt.Printf("true is : %v\n", true)
	}
	if !true {
		fmt.Printf("!true is : %v\n", !true)
	}
	if false {
		fmt.Printf("false is : %v\n", false)
	}
	if !false {
		fmt.Printf("!false is : %v\n", !false)
	}
	if 2 == 2 {
		fmt.Printf("2 == 2 is : %v\n", 2 == 2)
	}
	if 2 != 2 {
		fmt.Printf("2 != 2 is : %v\n", 2 != 2)
	}
	if !(2 == 2) {
		fmt.Printf("!(2 == 2) is : %v\n", !(2 == 2))
	}
	if !(2 != 2) {
		fmt.Printf("!(2 != 2) is : %v\n", !(2 != 2))
	}

	// Conditional - if statement with initialization statement
	fmt.Println("\n============= if statement with initialization statement =============")
	if ff := 23; ff <= 24 {
		fmt.Printf("%v <= 24\n", ff)
	}

	// Conditional - if statement - else if else if ... else
	fmt.Println("\n============= if statement - else if else if ... else =============")
	givenvalue := 77
	if givenvalue <= 10 {
		fmt.Printf("%v is between 0 - 10 range\n", givenvalue)
	} else if givenvalue <= 20 {
		fmt.Printf("%v is between 0 - 20 range\n", givenvalue)
	} else if givenvalue <= 30 {
		fmt.Printf("%v is between 0 - 30 range\n", givenvalue)
	} else if givenvalue <= 40 {
		fmt.Printf("%v is between 0 - 40 range\n", givenvalue)
	} else if givenvalue <= 50 {
		fmt.Printf("%v is between 0 - 50 range\n", givenvalue)
	} else {
		fmt.Printf("%v is not between 0 - 50 range\n", givenvalue)
	}

	// Conditional - Swicth statement
	fmt.Println("\n============= Swicth statement  =============")
	switch {
	case false:
		{
			fmt.Printf("Since the value of 'false' is %v, it's printing\n", false)
		}
	case true:
		{
			fmt.Printf("Since the value of 'true' is %v, it's printing\n", true)
		}
		fallthrough // if this block become true then execute next case also
	case 2 == 2:
		{
			fmt.Printf("Since the value of '2==2' is %v, it's printing\n", 2 == 2)
		}
	case 3 == 4:
		{
			fmt.Printf("Since the value of '3 == 4' is %v, it's printing\n", 3 == 4)
		}
	case 5 == 7:
		{
			fmt.Printf("Since the value of '5 == 7' is %v, it's printing\n", 5 == 7)
		}
	default:
		{
			fmt.Println("this is default")
		}
	}
}
