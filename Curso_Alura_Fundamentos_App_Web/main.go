package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	// 8000 é a porta, para testar seria localhost:8000
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39.00, Quantidade: 5},
		{"Tenis", "Confortável", 89.99, 3},
		{"Fone", "Muito bom", 59.50, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}
