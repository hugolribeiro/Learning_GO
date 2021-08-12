package main

import (
	"fmt"
	"reflect"
)

func main() {
	exibeNomes()
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.mercadolivre.com.br"
	fmt.Println(reflect.TypeOf(sites))

}

func exibeNomes() {
	nomes := []string{"Hugo", "Daniel", "Douglas", "Bernardo"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
	nomes = append(nomes, "Berenice")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
