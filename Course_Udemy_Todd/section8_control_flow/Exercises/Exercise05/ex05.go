package main

import "fmt"

/*
Instruction:
Print out the remainder (modulus) which is found for each number between 10 and 100
when it is divided by 4.
*/

func main() {
	for number := 10; number <= 100; number++ {
		fmt.Printf("When %d is divided by %d, the remainder (modulus) is %d\n", number, 4, number%4)
	}
}
