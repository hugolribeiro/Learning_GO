package main

import "fmt"

/*
Build and use an anonymous func
*/

func main() {
	numbers := []int{1, 2, 3, 4}
	x := func(n []int) int {
		var sum int
		for _, number := range n {
			sum += number
		}
		return sum
	}(numbers)
	fmt.Println(x)
}
