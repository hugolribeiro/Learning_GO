package main

import "fmt"

func main() {

	// Isso é uma função anônima. Podemos passar argumento normalmente e também podemos armazenar seu retorno em uma variável
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Argumento")

	fmt.Println(retorno)
}
