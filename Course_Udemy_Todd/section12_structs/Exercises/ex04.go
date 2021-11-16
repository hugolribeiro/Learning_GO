package main

import "fmt"

/*
Create and use an anonymous struct
*/

func main() {
	p1 := struct {
		firstName string
		lastName  string
	}{
		firstName: "Alan",
		lastName:  "Huan",
	}
	fmt.Println(p1)
}
