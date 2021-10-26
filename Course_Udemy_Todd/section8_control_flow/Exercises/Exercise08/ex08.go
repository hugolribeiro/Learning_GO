package main

import "fmt"

/*
Create a program that uses a switch statement with no switch expression specified
*/

func main() {
	switch {
	case 1 == 1:
		fmt.Println(1)
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("True")
	default:
		fmt.Println("The default")
	}
}
