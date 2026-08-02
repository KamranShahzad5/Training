package main

import (
	"Golang-Learning/utils"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	utils.PrintName()

	fmt.Println("Learning Variables")

	// Variable declaration and initialization
	var name string = "Kamran"
	fmt.Println(name)
	var age = 24

	fmt.Println(age)

	var rollno int = 23 // we can overwrite the value
	rollno = 25
	fmt.Println(rollno)

	var dimension float32 = 23.5
	fmt.Println(dimension)

	var decided bool = false
	fmt.Println(decided)

	const pi = 3.14 // value is fixed we canot over write it
	fmt.Println(pi)

	// Another way to declare a variable
	city := "New York"
	fmt.Println(city)

	// Zero Values

	var age1 int
	fmt.Println(age1)

	var name2 string
	fmt.Println(name2)

	var name3 bool
	fmt.Println(name3)

	// iota
	const (
		Red = iota
		Green
		Blue
	)

	fmt.Println(Red)
	fmt.Println(Green)
	fmt.Println(Blue)

	// Type Conversations

	var price float64 = 99.99
	fmt.Println("Before Converstion", price)

	var wholePrice int = int(price)

	fmt.Println("Upadte to int", wholePrice)

	str := "65"

	num, _ := strconv.Atoi(str)

	fmt.Println("Original string:", str)
	fmt.Println("Converted integer:", num)

}
