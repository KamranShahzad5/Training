package main

import "fmt"

func main() {
	var arr [5]string

	arr[0] = "Kamran"
	arr[1] = "Mango"
	arr[2] = "Apple"
	arr[3] = "Banana"
	arr[4] = "Orange"

	fmt.Println("Arr is ", arr)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Arr2 is ", arr2)

	fmt.Println("Length of arr is ", len(arr2))

}
