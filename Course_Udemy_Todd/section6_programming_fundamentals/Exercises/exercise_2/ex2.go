package main

import "fmt"

/* Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
*/

func main() {
	a := (42 == 42)
	b := (10 <= 40)
	c := (10 >= 40)
	d := ("a" != "a")
	f := (45.6 < 45.7)
	g := (45.7 > 45.6)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(g)
}
