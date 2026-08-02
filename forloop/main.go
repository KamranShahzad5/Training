package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop is", (i))

	}
	// for loop with slice
	slices := []int{1, 2, 3, 4, 5}
	for j := 0; j < 5; j++ {
		fmt.Println("Print slices values", slices[j])
	}

	// For loop for slice with range

	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Println("slice have index", index)
		fmt.Println("slice have value", value)

	}
	// For loop with strings

	username := "Kamran123"

	for index, ch := range username {

		fmt.Printf("Index: %d Character: %c\n", index, ch)
	}
}
