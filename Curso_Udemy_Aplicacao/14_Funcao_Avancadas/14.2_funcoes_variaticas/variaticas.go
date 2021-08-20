package main

import "fmt"

// A função receberá n números e os armazenará em uma slice (nesse caso "numeros")
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Podemos passar dois parâmetros em uma mesma função, sendo um normal e o outro variático.
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 10, 2, 3, 4, 5, 6, 7)
}
