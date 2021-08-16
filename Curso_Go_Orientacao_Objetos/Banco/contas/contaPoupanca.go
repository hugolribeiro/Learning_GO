package contas

import "Curso_Go_Orientacao_Objetos/Banco/clientes"

type ContaPoupanca struct {
	Titular                                clientes.Titular
	NumeroAgencia, NumeroDaConta, Operacao int
	saldo                                  float64
}

// Sacar Esse c *ContaCorrente  é o referente ao "this" em outras linguagens
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo Insuficiente"
	}
}

// Método getter de um atributo privado (letra inicial minúscula)
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

// Função com retorno múltiplo
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}
