package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	// Paralelismo depende da quantidade de núcleos do processador
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // quantidade de go routines a serem rodadas ao mesmo tempo

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // tira 1 do contador do waitGroup   -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // Vai esperar até chegar em 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
