package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 10, 42}
	x = append(x, 43, 44, 45)
	for index, value := range x {
		fmt.Println(index, value)
	}

	x = append(x[:2], x[4:]...)
	for index, value := range x {
		fmt.Println(index, value)
	}
}
