package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida dos bobos")
	fmt.Println(tipoEndereco)
}
