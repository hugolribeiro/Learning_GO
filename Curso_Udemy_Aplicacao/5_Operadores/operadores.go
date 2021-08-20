package main

import "fmt"

func main() {
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	// Como o Go é fortmente tipado, não podemos utilizar variáveis de diferentes tipos, então temos que usar um cast
	soma1 := numero1 + int16(numero2)
	fmt.Println(soma1)

	// FIM DOS ARITMÉTICOS

	// OPERADORES DE ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS - sempre retornam um booleano (true ou false)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("--------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	numero3 := 10
	numero3++
	numero3 += 5
	fmt.Println(numero3)

	numero3--
	numero3 -= 20
	fmt.Println(numero3)

	numero3 *= 3
	// FIM DOS OPERADORES UNÁRIOS

}
