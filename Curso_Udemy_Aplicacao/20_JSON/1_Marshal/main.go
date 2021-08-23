package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` //dessa forma especificamos dentro do struct qual chave se tornará dentro do JSON
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// usamos o Marshal para converter um map/struct para JSON
	c := cachorro{"Rex", "Dálmata", 3}

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON) // Isso trará apenas um slice de bytes

	// usamos o pacote bytes.NewBuffer para converter em buffer o slice em bytes trazido pelo Marshal
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	// Podemos também usar um map e convertê-lo para JSON
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
