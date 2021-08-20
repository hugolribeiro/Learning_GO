package main

import (
	"errors"
	"fmt"
)

func main() {
	// existem 4 tipos de números inteiros em go. A diferença é apenas a quantidade de bits que ele suporta
	// int8, int16, int32, int64        o "int" criará um int baseado na arquitetura do seu computador
	var numero int = 100000000000
	fmt.Println(numero)

	// uint aceita números inteiros SEM o indicador de negativo, ou seja, só aceitará números positivos
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// INT8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	// Tipos de números reais
	// float32, float64

	// Tipos string
	var str string = "Texto"
	fmt.Println(str)

	// O mais perto de char que temos em go
	char := 'B' //é o número que o char está na tabela ASCII
	fmt.Println(char)

	// FIM STRINGS

	// Valor zero

	var texto string
	fmt.Println(texto)

	var numerozero int
	fmt.Println(numerozero)

	// Booleano
	var booleano1 bool = true
	fmt.Println(booleano1)

	// Error
	var erro error
	fmt.Println(erro)

	// Erro usando o pacote errors
	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
