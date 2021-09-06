package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper() // O Helper fará com que o número da linha relatará em caso de erro seja o da linha de fora dessa função, ou seja, a linha real de erro
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz 'Olá' e o nome da pessoa", func(t *testing.T) {
		resultado := Ola("Hugo")
		esperado := "Olá, Hugo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, Mundo' quando uma string vazia for passada ", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
