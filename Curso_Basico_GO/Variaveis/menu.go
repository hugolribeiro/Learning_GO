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
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	fmt.Println("O endereço da minha variável comando é:", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
