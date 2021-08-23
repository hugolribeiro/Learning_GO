package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // Caso queiramos omitir algum campo, basta colocar `json:"-"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// Vamos converter um JSON para struct
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	fmt.Println(c)

	err := json.Unmarshal([]byte(cachorroEmJSON), &c) // Precisa receber um slice de bytes e também o ponteiro para o qual salvará as informações
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	// Vamos converter um json para map

	cachorro2EmJSON := `{"nome":"Toby", "raca":"Poodle"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
