package main

import "fmt"

/*
- Using a COMPOSITE LITERAL
	- create a SLICE of TYPE int
	- assign 10 values
- Range over the slice and print the values out
- Using format print
	- Print out the the type of slice
*/

func main() {
	numbers := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, number := range numbers {
		fmt.Println(number)
	}
	fmt.Printf("%T\n", numbers)
}
