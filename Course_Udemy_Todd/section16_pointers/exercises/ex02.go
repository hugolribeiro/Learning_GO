package main

import "fmt"

/*
- create a person struct
- create a func called "changeMe"with a *person as a parameter
	- change the value stored at the *person address
		- important
		- to dereference a struct, use (*value).field
		- As an exception, if the type of x is a named pointer type and (*x).f is a valid
	selector expression denoting a field (but not a method), x.f is shorthand for (*x).f
- create a value of type person
	- print out the value
- in func main
	- call "changeMe"
- in func main
	- print out the value

*/

type person struct {
	firstName string
	lastName  string
}

func changeMe(p1 *person) {
	p1.firstName = "Hugo"
}

func main() {
	p1 := person{
		firstName: "Lucas",
		lastName:  "Silva",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
