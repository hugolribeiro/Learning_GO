package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
