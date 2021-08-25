package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar Vai receber um router "cru" e retornar um router configurado (com todas as rotas)
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
