package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func Ola(name string) string {
	if name == "" {
		name = "Mundo"
	}
	return fmt.Sprintf("%s%s", prefixoOlaPortugues, name)
}

func main() {
	fmt.Println(Ola("mundo"))
}
