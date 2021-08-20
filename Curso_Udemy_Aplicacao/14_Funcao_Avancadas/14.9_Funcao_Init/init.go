package main

import "fmt"

// Uma função init é uma função que é executada antes da função main
// Podemos ter apenas 1 função init POR ARQUIVO
func init() {
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Função main sendo executada")
}

// output
// Executando a função init
// Função main sendo executada
