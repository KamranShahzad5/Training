package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "says Woof!")
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "says Meow!")
}

func MakeSpeak(s Speaker) {
	s.Speak()
}

func main() {
	dog := Dog{Name: "Tommy"}
	cat := Cat{Name: "Kitty"}

	MakeSpeak(dog)
	MakeSpeak(cat)
}