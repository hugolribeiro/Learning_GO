package main

import (
	"fmt"
	"math"
)

/*
create a type square
create a type circle
attach a method to each that calculates area and returns it
create a type shape which defines an interface as anything which has the area method
create a func info which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
user func info to print the area of circle
*/

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return s.side * s.side
}

func main() {
	c := circle{
		radius: 3,
	}
	s := square{
		side: 4,
	}
	info(c)
	info(s)
}
