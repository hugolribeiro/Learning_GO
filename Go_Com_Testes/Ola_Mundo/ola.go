package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"

const prefixoOlaFrances = "Bonjour, "
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "

func Ola(name string, idioma string) string {
	if name == "" {
		name = "Mundo"
	}
	if idioma == espanhol {
		return fmt.Sprintf("%s%s", prefixoOlaEspanhol, name)
	}
	if idioma == frances {
		return fmt.Sprintf("%s%s", prefixoOlaFrances, name)
	}
	return fmt.Sprintf("%s%s", prefixoOlaPortugues, name)
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
