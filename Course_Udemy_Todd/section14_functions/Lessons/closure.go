package main

import "fmt"

// Closure is when we have "enclosed" the scope of a variable in some code block.
var x int

func main() {
	var x int
	fmt.Println(x)
	x++

	// this is a closure block
	{
		y := 42
		fmt.Println(y)
	}
	// fmt.Println(y) // it will result in an error
	fmt.Println(x)
}
