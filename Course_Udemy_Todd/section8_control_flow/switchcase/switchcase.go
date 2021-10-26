package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough //this will force to continue in next case even though a false case
	case (4 == 4):
		fmt.Println("also true")
	default:
		fmt.Println("This is default")
	}

}
