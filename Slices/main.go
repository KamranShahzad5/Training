package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	slice = append(slice, 8, 9, 10, 11, 12, 13)
	fmt.Println("Slice have number", slice)
	fmt.Println("Slice have len", len(slice))

	// Slice 2.0 with make

	slices := make([]int, 3, 5)
	slices = append(slices, 4, 5, 6)
	fmt.Println("Slice have number", slices)
	fmt.Println("Slice have len", len(slices))
	fmt.Println("Slice have cap", cap(slices))

}
