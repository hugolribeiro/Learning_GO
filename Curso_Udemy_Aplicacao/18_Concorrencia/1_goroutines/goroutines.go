package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	// Paralelismo depende da quantidade de núcleos do processador
	go escrever("Olá Mundo!") // execute essa função, mas não espere o término para começar a executar a próxima (linha 12)
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
