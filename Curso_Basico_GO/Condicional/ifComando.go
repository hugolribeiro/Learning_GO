package main

import "fmt"

func main() {
	nome := "Hugo"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa")
	} else {
		fmt.Println("Não conheço este comando")
	}
}
