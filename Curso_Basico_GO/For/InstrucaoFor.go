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
	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.uol.com.br"}

	fmt.Println(sites)

	// for clássico
	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
	}

	// o for usando o range nos traz dois dados: o índice e o item em si
	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testaSite(site)
	}
	fmt.Println()

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}
