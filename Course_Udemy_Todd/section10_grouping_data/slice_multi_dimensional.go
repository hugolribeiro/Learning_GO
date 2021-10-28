package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(x)
	mp := []string{"Miss", "Moneypenny", "Strawberry"}
	fmt.Println(mp)

	y := [][]string{x, mp}
	fmt.Println(y)

}
