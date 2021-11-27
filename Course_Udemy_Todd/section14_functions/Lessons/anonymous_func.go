package main

import "fmt"

func main() {
	func() {
		fmt.Println("I am an anonymous func")
	}()
	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)
}
