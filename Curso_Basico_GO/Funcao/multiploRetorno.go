package main

import "fmt"

func main() {
	_, idade := devolveNomeEIdade()
	fmt.Println("Tenho ", idade, "anos")
}

func devolveNomeEIdade() (string, int) {
	nome := "Hugo"
	idade := 31
	return nome, idade
}
