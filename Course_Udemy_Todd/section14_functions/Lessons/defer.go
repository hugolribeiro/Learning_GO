package main

import "fmt"

func main() {
	defer foo()
	bar()
	teste()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func teste() {
	fmt.Println("teste")
}
