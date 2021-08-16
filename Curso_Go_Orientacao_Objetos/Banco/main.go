package main

import (
	"Curso_Go_Orientacao_Objetos/Banco/contas"
	"fmt"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo)

	fmt.Println(contaExemplo.ObterSaldo())

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(150)
	contaDoDenis.ObterSaldo()

	fmt.Println(contaDoDenis)

}

//func TestandoObjetos() {
//	// Apenas testando criações e comparações de objetos
//	fmt.Println("---------------- TESTES ---------------------")
//	// Primeira maneira de se instanciar um objeto
//	contaDoGuilherme := contas.ContaCorrente{
//		Titular:       "Guilherme",
//		NumeroAgencia: 589,
//		NumeroConta:   123456,
//		Saldo:         125.5,
//	}
//
//	contaDaBruna := contas.ContaCorrente{
//		Titular:       "Bruna",
//		NumeroAgencia: 222,
//		NumeroConta:   111222,
//		Saldo:         200,
//	}
//
//	// Segunda maneira de se instanciar um objeto - Mais utilizada
//	contaDoHugo := contas.ContaCorrente{"Hugo", 333, 111333, 300}
//
//	fmt.Println(contaDoGuilherme)
//	fmt.Println(contaDaBruna)
//	fmt.Println(contaDoHugo)
//
//	// Terceira maneira de se instanciar um Objeto
//	var contaDaCris *contas.ContaCorrente
//	contaDaCris = new(contas.ContaCorrente)
//	contaDaCris.Titular = "Cris"
//	contaDaCris.Saldo = 500
//
//	// Comparando contas
//	fmt.Println("Comparando contas")
//	contaJoao1 := contas.ContaCorrente{"Joao", 444, 111444, 400}
//	contaJoao2 := contas.ContaCorrente{"Joao", 444, 111444, 400}
//
//	fmt.Println(contaJoao1 == contaJoao2) // Resultado será true
//
//	contaDaBruna2 := contas.ContaCorrente{"Bruna", 222, 111222, 200}
//
//	fmt.Println(contaDaBruna == contaDaBruna2) // Resultado será true
//
//	var contaDaCris2 *contas.ContaCorrente
//	contaDaCris2 = new(contas.ContaCorrente)
//	contaDaCris2.Titular = "Cris"
//	contaDaCris2.Saldo = 500
//
//	fmt.Println(contaDaCris == contaDaCris2)   // Resultado será false! Pois os endereços na memória são diferentes
//	fmt.Println(*contaDaCris == *contaDaCris2) // O resultado será true, pois agora estamos comparando o conteúdo

//}
