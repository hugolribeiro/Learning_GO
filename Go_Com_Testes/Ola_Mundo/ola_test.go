package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Hugo")
	esperado := "Olá, Hugo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
