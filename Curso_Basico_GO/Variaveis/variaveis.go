package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Hugo"
	idade := 31
	versao := 1.1
	var andando bool = true
	fmt.Println("Olá", nome, ". Sua idade é :", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é:", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versão é:", reflect.TypeOf(versao))
	fmt.Println("O tipo da variável andando é:", reflect.TypeOf(andando))
}
