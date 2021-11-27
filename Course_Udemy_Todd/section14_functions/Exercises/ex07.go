package main

import "fmt"

/*
Assign a func to a variable, then call that func
*/

func main() {
	sumFunc := func(numbers []int) int {
		var sum int
		for _, number := range numbers {
			sum += number
		}
		return sum
	}
	sum := sumFunc([]int{1, 2, 3})
	fmt.Println(sum)
}
