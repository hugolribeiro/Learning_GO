package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const alemao = "alemao"

const prefixoOlaFrances = "Bonjour, "
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaAlemao = "Hallo, "

func Ola(name string, idioma string) string {
	if name == "" {
		name = "Mundo"
	}

	return fmt.Sprintf("%s%s", prefixoDeSaudacao(idioma), name)

}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case alemao:
		prefixo = prefixoOlaAlemao
	default:
		prefixo = prefixoOlaPortugues
	}
	return prefixo
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
