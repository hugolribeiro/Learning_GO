package main

import (
	"fmt"
)

/*
Using a COMPOSITE LITERAL
	- Create an ARRAY which hold 5 values of type int
	- assign VALUES to each index position
Range over the array and print the values out
Using format print
	- print out the TYPE of the array
*/

func main() {
	var numbers [5]int
	for i := 0; i < 5; i++ {
		numbers[i] = i + 10
	}
	for _, number := range numbers {
		fmt.Println(number)
	}
	fmt.Printf("%T", numbers)
}
