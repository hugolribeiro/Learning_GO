package main

import "fmt"

// O canal com buffer somente ficar[a bloqueado ao atingir a capacidade demonstrada
func main() {
	canal := make(chan string, 2) // especificando que o canal possui capacidade de 2
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
