package main

import "fmt"

/*
Take the code from the previous exercise, then store the values of type person
in a map with the key of last name.
Access each value in the map. Print out the values, ranging over the slice
*/

type person struct {
	firsName    string
	lastName    string
	favIceCream []string
}

func main() {
	p1 := person{
		firsName:    "Carlos",
		lastName:    "Navarro",
		favIceCream: []string{"Chocolate", "Strawberry"},
	}
	p2 := person{
		firsName:    "James",
		lastName:    "Ferrera",
		favIceCream: []string{"Pineable", "Watermelon"},
	}
	people := make(map[string]person)
	people["Navarro"] = p1
	people["Ferrera"] = p2
	fmt.Println(people["Navarro"])
	fmt.Println(people["Ferrera"])
}
