package main

import "fmt"

func main() {
	// Primeira maneira de declaração
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// Segunda maneira (mais usada) -  o nome técnico é inferência de tipo
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	// Terceira maneira
	var (
		variavel3 string = "lalala"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3, variavel4)

	// Quarta maneira
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	// Declaração de constante
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

}
