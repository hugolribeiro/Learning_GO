package main

import "fmt"

/*
Closure is when we have "enclosed" the scope of a variable in some code block.
Create a func which "encloses"the scope of a variable
*/

func main() {
	x := 50
	showNumber()
	fmt.Println(x)
}

func showNumber() {
	x := 42
	fmt.Println(x)
}
