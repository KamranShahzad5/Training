package main

import "fmt"

type Person struct{
	name string
	age int
	phoneNo int
}

func (p Person) getdata(){
	fmt.Println("my name is",p.name)
	fmt.Println("my age is",p.age)
	fmt.Println("my phone number is",p.phoneNo)
}

func main (){
	person:=Person{
		name:"Kamran",
		age:22,
		phoneNo:12345,
	}
	person.getdata()
}