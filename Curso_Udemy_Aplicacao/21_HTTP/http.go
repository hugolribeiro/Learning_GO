package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// REQUEST - RESPONSE

	// Rotas = maneira de identificar o tipo de mensagem que está sendo enviada e que tipo de providência tomar
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página Raiz!"))
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuários"))
	})

	log.Fatal(http.ListenAndServe(":8000", nil)) // com essa linha já teremos um servidor rodando na porta especificada

}
