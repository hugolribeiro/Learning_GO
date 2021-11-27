package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	// That way we unfurly a slice
	x := sum(xi...)
	fmt.Println("The total is: ", x)
}

func sum(x ...int) int {
	sum := 0
	for _, number := range x {
		sum += number
	}
	return sum
}
