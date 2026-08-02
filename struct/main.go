package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Details struct {
	phone int
	email string
}

// Struct 3 (contains other structs)
type Employee struct {
	person  Person
	details Details
	role    string
	salary  int
}

func main() {

	// Method 1:
	var myType Person
	myType.name = "Alice"
	myType.age = 30

	fmt.Println(myType)

	// Method 2:
	myType2 := Person{
		name: "Bob",
		age:  25,
	}

	fmt.Println(myType2)

	// Create Employee
	employee := Employee{
		person: Person{
			name: "John",
			age:  28,
		},
		details: Details{
			phone: 123456789,
			email: "john@gmail.com",
		},
		role:   "Software Engineer",
		salary: 120000,
	}

	fmt.Println(employee)

	fmt.Println("Name:", employee.person.name)
	fmt.Println("Age:", employee.person.age)
	fmt.Println("Phone:", employee.details.phone)
	fmt.Println("Email:", employee.details.email)
	fmt.Println("Role:", employee.role)
	fmt.Println("Salary:", employee.salary)
}
