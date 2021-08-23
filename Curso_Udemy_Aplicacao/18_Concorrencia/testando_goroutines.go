package main

import (
	"fmt"
	"time"
)

func main() {

	go func() { fmt.Println("oi sou uma goroutine1!") }()
	go func() { fmt.Println("oi sou uma goroutine2!") }()
	go func() { fmt.Println("oi sou uma goroutine3!") }()
	fmt.Println("Olá, estamos testando as goroutines em Go!")
	time.Sleep(time.Second)

}
