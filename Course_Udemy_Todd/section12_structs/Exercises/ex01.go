package main

import "fmt"

/*
Create your own type "person" which will have an underlying type of "struct"
so that it can store the following data:
	- first name
	- last name
	- favorite ice cream flavors
Create two Values of Type person. Print out the values, ranging over the elements in the slice.
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
	fmt.Println(p1.firsName)
	fmt.Println(p1.lastName)
	for _, favlor := range p1.favIceCream {
		fmt.Println(favlor)
	}

	fmt.Println(p2.firsName)
	fmt.Println(p2.lastName)
	for _, favlor := range p2.favIceCream {
		fmt.Println(favlor)
	}
}
