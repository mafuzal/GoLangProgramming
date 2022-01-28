package main

import "fmt"

var y = 3.71

// Declare a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 fir integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps

var z int

func main() {
	//Short declaration operator
	//Declare a variable and ASSIGN a VALUE (of a certian type)
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	Mars()
	fmt.Println(z)
}

func Mars() {
	fmt.Println("Gravity on Mars ", y, "m/s Square")
}
