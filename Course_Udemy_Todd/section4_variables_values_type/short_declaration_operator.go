package main

import "fmt"

// global scope
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" if of TYPE int

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
