package main

import "fmt"

/*
Using the code from the previous example delete a record from your map. Now print the map out using the "range" loop.
*/

func main() {
	persons := map[string][]string{}
	persons["Joao"] = []string{"cook", "ride bike"}
	persons["Maria"] = []string{"hunt witches", "eat breads"}
	persons["Clodoaldo"] = []string{"swim", "eat pizza"}

	delete(persons, "Maria")

	for key, value := range persons {
		fmt.Println("This is the record for", key)
		for i, v2 := range value {
			fmt.Println("\t", i, v2)
		}
	}
}
