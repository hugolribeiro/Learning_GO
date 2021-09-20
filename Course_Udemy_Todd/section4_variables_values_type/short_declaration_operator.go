package main

import "fmt"

// global scope
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" if of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	fmt.Println(z)
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	fmt.Println(y)
	y = 100 + 70
	fmt.Println(y)
	w := "Bond, James"
	fmt.Println(w)
}
