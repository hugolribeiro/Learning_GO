package main

import "fmt"

// retornará uma func()
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}

// Outout:
// Dentro da função main
// Dentro da função closure
