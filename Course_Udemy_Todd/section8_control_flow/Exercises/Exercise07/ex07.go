package main

import "fmt"

/*
Building on the previoues hand-on exercise, create a program that uses "else if" and "else"
*/

func main() {
	var age int
	fmt.Println("Input your age: ")
	fmt.Scan(&age)
	if age < 18 {
		fmt.Println("You are lesser than 18 years old")
	} else if age == 18 {
		fmt.Println("You are exactly 18 years old")
	} else {
		fmt.Println("You are over than 18 years old")
	}
}
