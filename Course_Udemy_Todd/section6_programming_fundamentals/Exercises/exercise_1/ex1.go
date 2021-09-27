package main

import "fmt"

// Write a program that prints a number in decimal, binary, and hex

func main() {
	number := 42
	fmt.Printf("42 in decimal: %d\n", number)
	fmt.Printf("42 in binary: %#b\n", number)
	fmt.Printf("42 in hexadecimal: %#x\n", number)
}
