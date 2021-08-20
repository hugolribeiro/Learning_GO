package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if init (quando executa uma condição if e já inicia uma variável)
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número menor ou igual a zero")
	}

	// não é possível acessar a variável "outroNumero" fora do escopo do if
	// fmt.Println(outroNumero) - isso não deixará compilar
}
