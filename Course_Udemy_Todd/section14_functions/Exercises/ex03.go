package main

import "fmt"

/*
Use Defer
*/

func main() {
	defer foo()
	fmt.Println("Hello")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}
