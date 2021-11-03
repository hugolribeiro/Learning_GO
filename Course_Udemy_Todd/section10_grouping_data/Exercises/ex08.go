package main

import "fmt"

/*
Create a map with a key of TYPE string which is a person's last_first name, and a value of type
[]string which stores their favorite things. Store three records in your map.
Print out all of the values, along with their index position in the slice.
*/

func main() {
	persons := map[string][]string{}
	persons["Joao"] = []string{"cook", "ride bike"}
	persons["Maria"] = []string{"hunt witches", "eat breads"}
	persons["Clodoaldo"] = []string{"swim", "eat pizza"}

	for key, value := range persons {
		fmt.Println("This is the record for", key)
		for i, v2 := range value {
			fmt.Println("\t", i, v2)
		}
	}
}
