package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// ARRAYS POSSUEM TAMANHO E TIPOS DE DADOS FIXOS
	//Primeira maneira para se criar um array
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	// Segunda maneira para se criar um array
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posicão 4", "Posição 5"}
	fmt.Println(array2)

	// Terceira maneira para se criar um array
	// com os três pontos estamos dizendo que o tamanho dele será a quantidade de valores que passarmos entre {}
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICE - estrutura baseada no array, mas com tamanho dinâmico
	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)
	slice = append(slice, 17) // adicionando elemento ao slice
	fmt.Println(slice)

	slice2 := array2[1:3] // o primeiro índice é INCLUSIVO e o segundo índice é EXCLUSIVO
	fmt.Println(slice2)

	array2[1] = "Posiçäo Alterada"
	fmt.Println(slice2)

	// Vendo o tipo de cada variável
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// Arrays Internos
	// A função make exige 3 parâmetros: o tipo, tamanho e capacidade
	fmt.Println("-------- Arrays Internos ---------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho atual
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 10) // quando estouramos a capacidade do slice, o Go dobrará a capacidade
	fmt.Println(cap(slice3))

	// podemos criar um slice sem declarar a capacidade, assim, a capacidade será igual ao tamanho dele
	slice4 := make([]float32, 5)
	fmt.Println(cap(slice4))
}
