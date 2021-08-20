package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// funções podem ter mais de 1 retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// declarando uma variável como uma função
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	var resultado string
	resultado = f("Texto da funcao1")
	fmt.Println(resultado)

	// duas variáveis recebendo os resultados de uma função
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Podemos suprimir alguma variável que não usaremo, basta inserir "_"
	resultadoSoma1, _ := calculosMatematicos(20, 30)
	fmt.Println(resultadoSoma1)
}
