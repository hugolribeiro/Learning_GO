package main

import "fmt"

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
- "James", "Bond", "Shaken, not stirred"
- "Miss", "moneypenny", "Helloooooo, James".
Range over the records, then range over the data in each record.
*/

func main() {
	x := [][]string{
		{
			"James", "Bond", "Shaken, not stirred",
		},
		{
			"Miss", "moneypenny", "Hellooooo, James",
		},
	}
	for i, x1 := range x {
		fmt.Println("record: ", i)
		for j, v := range x1 {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, v)
		}
	}
}
