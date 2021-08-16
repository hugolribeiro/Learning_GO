package contas

import "Curso_Go_Orientacao_Objetos/Banco/clientes"

// Uma estrutura se assemelha ao conceito de Classe em POO

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

// Sacar Esse c *ContaCorrente  é o referente ao "this" em outras linguagens
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo Insuficiente"
	}
}

// Função com retorno múltiplo
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

// Detalhe para o asterisco em ContaCorrente para indicar a referência à conta de destino
func (contaOrigem *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < contaOrigem.saldo && valorDaTransferencia > 0 {
		contaOrigem.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

// Método getter de um atributo privado (letra inicial minúscula)
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
