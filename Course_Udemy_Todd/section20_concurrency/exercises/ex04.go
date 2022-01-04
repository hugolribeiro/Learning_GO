/*
Fix the race condition you created in the program 03 by using a mutex
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	goRoutines := 100
	wg.Add(goRoutines)
	var m sync.Mutex

	for i := 0; i < goRoutines; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
