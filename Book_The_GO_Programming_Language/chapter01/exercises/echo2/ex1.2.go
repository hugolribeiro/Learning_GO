package main

import (
	"fmt"
	"os"
)

/*
Modify the Echo program to print the index and the value of each of its args, one per line.
*/

func main() {
	for i, v := range os.Args[:] {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}
}
