package main

import "fmt"

// Using iota create 4 constants for the last 4 years. Print the constant values

const (
	year4 = iota + 2017
	year3
	year2
	year1
)

func main() {
	fmt.Println(year4)
	fmt.Println(year3)
	fmt.Println(year2)
	fmt.Println(year1)
}
