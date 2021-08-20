package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Criando um map
	// dentro dos colchetes ficará o tipo da key do map
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"])

	// maps aninhados
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus1",
		},
	}

	fmt.Println(usuario2)

	// Apagando uma chave do map
	delete(usuario2, "nome")

	fmt.Println(usuario2)

	// Adicionando outra chave
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)

}
