package main

import "fmt"

/*
Create a variable of type string using raw string literal. Print it
*/

func main() {
	phrase := `This is a phrase
	using raw string
	we will print it`

	fmt.Println(phrase)
}
