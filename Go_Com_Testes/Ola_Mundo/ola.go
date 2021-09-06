package main

import "fmt"

func Ola(name string) string {
	return fmt.Sprintf("Olá, %s", name)
}

func main() {
	fmt.Println(Ola("mundo"))
}
