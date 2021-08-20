package main

import "fmt"

// Uma interface assim receberá qualquer interface com parâmetros de qualquer tipo
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
}

// Uma função que usa de tipos genéricos é justamente a Println, já que recebe diversos tipos de variáveis
