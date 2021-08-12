package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeintroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0) // informa para o SO que vc quer sair do programa (código 0)s
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) // informa para o SO que ocorreu algum problema inesperado no programa
		}
	}

}

func exibeintroducao() {
	nome := "Hugo"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Este programa está na versão", versao)

}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.mercadolivre.com.br"
	fmt.Println(sites[3])

	fmt.Println(sites)
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "esta com problemas. Status Code:", resp.StatusCode)
	}

}
