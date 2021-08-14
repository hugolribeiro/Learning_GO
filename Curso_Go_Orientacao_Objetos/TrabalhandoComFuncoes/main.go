package main

import (
	"fmt"
)

func SemParametro() string {
	return "Exemplo de função sem parâmetro"
}

func UmParametro(texto string) string {
	return texto
}

func DoisParametrosDoisRetornos(texto string, numero int) (string, int) {
	return texto, numero
}

// Criando uma função variádica (recebe um número indeterminado de parâmetros)
func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

func main() {
	fmt.Println(SemParametro())
	fmt.Println(UmParametro("Exemplo de função com um parâmetro"))
	fmt.Println(DoisParametrosDoisRetornos("Passando dois parâmetros: essa string e o número", 10))

	fmt.Println(Somando(1, 2))
	fmt.Println(Somando(1, 2, 3))
	fmt.Println(Somando(1, 2, 3, 4))

}
