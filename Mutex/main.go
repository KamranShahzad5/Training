package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex
var wg sync.WaitGroup

func increment() {
	defer wg.Done()

	mu.Lock() // Lock shared resource
	counter++
	mu.Unlock() // Unlock shared resource
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()

	fmt.Println(counter)
}
