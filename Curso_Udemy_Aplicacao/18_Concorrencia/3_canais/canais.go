package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) // só poderá enviar e receber strings
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// Primeira opção
	//for {
	//	mensagem, aberto := <- canal // Aqui estou rebendo uma mensagem enviada pelo canal
	//	if !aberto { // quando o canal estiver fechado, pararei o loop
	//		break
	//	}
	//
	//}

	// Segunda opção
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // canal estará recebendo esse texto ... estou mandando um valor para o canal
		time.Sleep(time.Second)
	}
	close(canal)
}
