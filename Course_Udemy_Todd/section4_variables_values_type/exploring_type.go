package main

import "fmt"

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I am ASSIGN the VALUE "Shaken, not stirred"

var z string = "Shaken, not stirred"
var a string = `James said, "Shaken, not stirred"`

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
