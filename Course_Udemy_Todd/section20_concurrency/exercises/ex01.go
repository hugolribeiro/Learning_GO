/*
In addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(5)

	fmt.Println("Helloooo")

	go func() {
		fmt.Println("Helllooo 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Helllooo 2")
		wg.Done()
	}()

	go func() {
		fmt.Println("Helllooo 3")
		wg.Done()
	}()

	go func() {
		fmt.Println("Helllooo 4")
		wg.Done()
	}()

	go func() {
		fmt.Println("Helllooo 5")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Exit program")
}
