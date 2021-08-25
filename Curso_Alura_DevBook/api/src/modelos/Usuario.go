package modelos

import "time"

// Usuario representa um usuário que utilizará a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"` // se for passar um usuário para json e não informar id, esse campo ficará vazio e não zero
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}
