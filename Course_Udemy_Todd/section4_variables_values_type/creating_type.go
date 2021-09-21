package main

import "fmt"

var a int

type hotdog int

var b hotdog

type abracadabra hotdog

var c abracadabra

func main() {
	a = 42
	b = 43
	c = 44
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	// convertion
	a = int(b)
	fmt.Println(a)
}
