package main

import "fmt"

/*
start with this slice:
- [42, 43, 44, 45, 46, 47, 48, 49, 50, 51]
append to that slice the value 52
print the slice
in ONE STATEMENT append to that slice these values
- 53, 54, 55
print the slice
append to the slice this slice
- [56, 57, 58, 59, 60]
print the slice
*/

func main() {
	mainSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	mainSlice = append(mainSlice, 52)
	fmt.Println(mainSlice)
	mainSlice = append(mainSlice, 53, 54, 55)
	fmt.Println(mainSlice)
	mainSlice = append(mainSlice, []int{56, 57, 58, 59, 60}...)
	fmt.Println(mainSlice)
}
