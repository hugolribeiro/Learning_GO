package main

import "fmt"

//métodos estão associados a algo, diferentemente das funções

type usuario struct {
	nome  string
	idade uint8
}

// esse é um método da struct usuário
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

// método que altera alguma informação de nossa struct. Devemos passar como um ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorIdade()
	fmt.Println(maiorDeIdade)

	fmt.Println(usuario2.idade)
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}

// output
// Salvando os dados do Usuário Usuário 1 no banco de dados
// Salvando os dados do Usuário Davi no banco de dados
// true
// 30
// 31
