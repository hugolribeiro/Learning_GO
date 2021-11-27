package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first func expression")
	}
	f2 := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	f()
	f2(1984)
}
