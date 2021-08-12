package main

import "fmt"

// Uma estrutura se assemelha ao conceito de Classe em POO

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
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

	contaDoHugo := ContaCorrente{"Hugo", 333, 111333, 300}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDoHugo)
}
