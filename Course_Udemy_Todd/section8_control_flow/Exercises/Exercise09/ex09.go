package main

import "fmt"

/*
Create a program that uses a switch statement with the switch expression specified as a
variable of TYPE string with the IDENTIFIER "favSport"
*/

func main() {
	favSport := "beach tennis"
	switch favSport {
	case "swimming":
		fmt.Println(1)
	case "beach tennis":
		fmt.Println("OKKKKK")
	default:
		fmt.Println("No matches")
	}
}
