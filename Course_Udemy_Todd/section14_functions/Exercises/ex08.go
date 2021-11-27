package main

import "fmt"

/*
Create a func which returns a func
assign the returned func to a variable
call the returned func
*/

func main() {
	returnedFunc := func() func([]int) int {
		return func(numbers []int) int {
			var sum int
			for _, number := range numbers {
				sum += number
			}
			return sum
		}
	}
	sumFunc := returnedFunc()
	sum := sumFunc([]int{1, 5, 10})
	fmt.Println(sum)
}
