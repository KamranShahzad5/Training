package main

import "fmt"

func main (){
	num :=2

	ptr:=&num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Value at address stored in ptr:", *ptr)
}
