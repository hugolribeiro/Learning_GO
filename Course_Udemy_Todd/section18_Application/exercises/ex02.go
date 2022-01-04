// Unmarshal the JSON into a Go data structure

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	json1 := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`
	fmt.Println(json1)

	var users []user

	err := json.Unmarshal([]byte(json1), &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)

}
