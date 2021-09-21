// Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS "x", "y" and "z"
// a 42
// b "James Bond"
// c true

package main

import "fmt"

func main() {
	a := 42
	b := "James Bond"
	c := true

	// Now, print the values stored in those variables using
	// a single print statement
	// b multiple print statements
	fmt.Printf("a: %d\nb: %s\nc: %t\n", a, b, c)
	fmt.Println("-------------")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}
