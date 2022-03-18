package main

import (
	"fmt"
	"os"
	"strings"
)

/*
	Modify the echo program to also print os.Args[0], that is the command name that called it.

	Solution: We can take the example of echo3 and pass the entire os.Args
*/

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
