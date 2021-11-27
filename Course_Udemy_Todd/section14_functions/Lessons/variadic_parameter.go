package main

import "fmt"

func main() {
	fmt.Println(foo(3, 4, 5, 6, 7))
}

// This is a variadic parameter
func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, number := range x {
		sum += number
	}
	return sum
}

// func (r receiver) identifier(parameter(s)) (return(s)) { code }
