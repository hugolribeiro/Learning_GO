package main

import "fmt"

// Nessa função estamos passando uma cópia daquela variável
func inverterSinal(numero int) int {
	return numero * -1
}

// Dessa maneira estamos passando uma referência àquela variável. Então alteramos de fato o valor da variável
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}

//output
// -20
// 20
// 40
// -40
