package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Hugo", 31, enderecoExemplo}
	fmt.Println(usuario2)

	// Caso queira criar um usuário sem possuir todos os dados ainda
	usuario3 := usuario{nome: "Alex"}
	fmt.Println(usuario3)
}
