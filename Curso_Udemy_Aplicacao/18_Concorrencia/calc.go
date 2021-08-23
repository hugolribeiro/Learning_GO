package main

import (
	"fmt"
)

func main() {

	go func() { fmt.Println(soma(5, 2)) }()
	go func() { fmt.Println(subtracao(5, 2)) }()
	go func() { fmt.Println(mult(3, 2)) }()
	go func() { fmt.Println(div(10, 2)) }()
	fmt.Println("Operações")
}

func div(i int, i2 int) int {
	return i / i2
}

func mult(i int, i2 int) int {
	return i * i2
}

func subtracao(i int, i2 int) int {
	return i - i2
}

func soma(i int, i2 int) int {
	return i + i2
}
