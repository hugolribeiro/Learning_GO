/*
This exercise will reinforce our understanding of method sets
- create a type person struct
- attach a method speak to type pointer to a person
	- *person
- create a type human interface
	- to implicity implement the interface, a human must have the speak method
- create func "saySomething"
	- have it take in a human as parameter
	- have it call the speak method
- show the following in your code
	- you CAN pass type *person into saySomething
	- you CANNOT pass type person into saySomething
*/

package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Printf("%s is speaking\n", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p1 := person{
		name: "Hugo",
	}
	saySomething(&p1)
}
