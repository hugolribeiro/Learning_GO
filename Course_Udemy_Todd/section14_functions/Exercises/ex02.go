package main

import "fmt"

/*
- create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
- create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
*/

func main() {
	sum := foo(1, 2, 3, 4)
	fmt.Println(sum)
	sum2 := bar([]int{1, 2, 3})
	fmt.Println(sum2)
}

func foo(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func bar(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
