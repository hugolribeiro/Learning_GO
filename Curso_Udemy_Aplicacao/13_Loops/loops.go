package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Inrementando i", i)
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	// Com essa declaração, a variável j estará apenas dentro do escopo do loop
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	// Usando loop para percorrer uma coleção. O "range" retornará o índice e o nome
	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// Podemos iterar por letra em uma string também, porém, retornará o índice e números daquela letra na tabela ASCII
	// Temos que usar o cast string() para converter da tabela ASCII para alfanumerico

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// Iterando em maps
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// // Loop infinito
	// for {
	// 	fmt.Println("Executando infinitamente")
	// 	time.Sleep(time.Second)
	// }

}
