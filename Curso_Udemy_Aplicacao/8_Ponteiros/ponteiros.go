package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1 // passamos uma CÓPIA do valor que está na variável1

	fmt.Println(variavel1, variavel2)

	variavel1++ // muda apenas o valor da variável 1
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro) // será  0   e   nil

	variavel3 = 100
	ponteiro = &variavel3 // O "&" vai pegar o endereço de memória em que está a variável3

	fmt.Println(variavel3, ponteiro) // 100 0xc0000ac008

	fmt.Println(variavel3, *ponteiro) // desreferenciação da variável ponteiro, assim ele lerá o valor que está no endereço de memória  | 100 100

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // 150  150
}
