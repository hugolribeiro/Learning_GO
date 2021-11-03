package main

import "fmt"

/*
Create a slice to store the names of all of the states in Brazil.
What is the length of your slice?
What is the capacity?
Print out all of the values, along with their index position in the slice, without using the range clause.
*/

func main() {
	states := []string{
		"Acre",
		"Alagoas",
		"Amapá",
		"Amazonas",
		"Bahia",
		"Ceará",
		"Espírito Santo",
		"Goiás",
		"Maranhão",
		"Mato Grosso",
		"Mato Grosso do Sul",
		"Minas Gerais",
		"Pará",
		"Paraíba",
		"Paraná",
		"Pernambuco",
		"Piauí",
		"Rio de Janeiro",
		"Rio Grande do Norte",
		"Rio Grande do Sul",
		"rondônia",
		"Roraima",
		"Santa Catarina",
		"São Paulo",
		"Sergipe",
		"Tocantins",
		"Distrito Federal",
	}
	length := len(states)
	fmt.Printf("Length: %d\n", length)
	fmt.Printf("Capacity: %d\n", cap(states))
	for i := 0; i < length; i++ {
		fmt.Println(states[i])
	}
}
