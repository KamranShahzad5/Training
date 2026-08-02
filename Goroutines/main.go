package main

import (
	"fmt"
	"time"
)

func sayhello() {
	fmt.Println("Kamran runs Goroutines")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Goroutines Continues")
}

func done(){
	fmt.Println("Goroutines Done")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("Hello, World!")
	go sayhello()
	go done()

	// Wait for the goroutine to finish
	time.Sleep(3000 * time.Millisecond)
	
}
