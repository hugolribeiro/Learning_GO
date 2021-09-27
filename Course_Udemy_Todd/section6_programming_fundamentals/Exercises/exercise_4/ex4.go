package main

import "fmt"

// assigns an int to a variable
// prints that int in decimal, binary, and hex
// shifts the bits of that int over 1 position to the left, and assigns that to a variable
// prints that variable in decimal, binary, and hex

func main() {
	number1 := 15
	fmt.Printf("number1 in decimal: %d\nnumber1 in binary: %#b\nnumber1 in hex: %#x\n", number1, number1, number1)
	number2 := number1 << 1
	fmt.Printf("number2 in decimal: %d\nnumber2 in binary: %#b\nnumber2 in hex: %#x\n", number2, number2, number2)

}
