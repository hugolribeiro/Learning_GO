package main

import "fmt"

/*
pass a func into a func as an argument.
*/

func sum(numbers []int) int {
	var sumNumbers int
	for _, number := range numbers {
		sumNumbers += number
	}
	return sumNumbers
}

func sumEven(f func(numbers []int) int, n []int) int {
	var evenNumbers []int
	for _, number := range n {
		if number%2 != 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}
	return f(evenNumbers)
}

func main() {
	allNumbers := []int{1, 2, 3, 4, 5}
	sumEvenNumbers := sumEven(sum, allNumbers)
	fmt.Println(sumEvenNumbers)
}
