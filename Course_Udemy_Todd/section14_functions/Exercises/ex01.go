package main

import "fmt"

/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 10, "Heyyy"
}
