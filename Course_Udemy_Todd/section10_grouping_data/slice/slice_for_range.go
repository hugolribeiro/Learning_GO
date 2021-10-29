package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}

	for index, value := range x {
		fmt.Println(index, value)
	}
}
