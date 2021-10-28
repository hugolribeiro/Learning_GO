package main

import "fmt"

/*
An array is a numbered sequence of elements of a single type, called the element type.
The number of elements is called the length and is never negative.
*/

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
