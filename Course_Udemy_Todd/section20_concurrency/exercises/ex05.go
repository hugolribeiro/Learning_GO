/*
Fix the race condition you created in exercise #4 by using package atomic
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64

	goRoutines := 100
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
