package main

import (
	"crud/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE

	// CREATE - POST
	// READ - GET
	// UPDATE - PUT
	// DELETE - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	fmt.Println("Escutando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
