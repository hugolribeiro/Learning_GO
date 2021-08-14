package main

import "fmt"

// Uma estrutura se assemelha ao conceito de Classe em POO

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Sacar Esse c *ContaCorrente  é o referente ao "this" em outras linguagens
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia)

	valorDoSaque := 200.0
	fmt.Println(contaDaSilvia.Sacar(valorDoSaque))

	valorDoSaque = 800
	fmt.Println(contaDaSilvia.Sacar(valorDoSaque))
}

func TestandoObjetos() {
	// Apenas testando criações e comparações de objetos
	fmt.Println("---------------- TESTES ---------------------")
	// Primeira maneira de se instanciar um objeto
	contaDoGuilherme := ContaCorrente{
		titular:       "Guilherme",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	contaDaBruna := ContaCorrente{
		titular:       "Bruna",
		numeroAgencia: 222,
		numeroConta:   111222,
		saldo:         200,
	}

	// Segunda maneira de se instanciar um objeto - Mais utilizada
	contaDoHugo := ContaCorrente{"Hugo", 333, 111333, 300}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDoHugo)

	// Terceira maneira de se instanciar um Objeto
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	// Comparando contas
	fmt.Println("Comparando contas")
	contaJoao1 := ContaCorrente{"Joao", 444, 111444, 400}
	contaJoao2 := ContaCorrente{"Joao", 444, 111444, 400}

	fmt.Println(contaJoao1 == contaJoao2) // Resultado será true

	contaDaBruna2 := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDaBruna == contaDaBruna2) // Resultado será true

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	fmt.Println(contaDaCris == contaDaCris2)   // Resultado será false! Pois os endereços na memória são diferentes
	fmt.Println(*contaDaCris == *contaDaCris2) // O resultado será true, pois agora estamos comparando o conteúdo

}
