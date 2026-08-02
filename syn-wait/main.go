package main

import (
	"fmt"
	"sync"
)

func workers(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease counter when goroutine finishes

	fmt.Println("Worker", i, "is working")
	fmt.Println("Worker", i, "is done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increase counter
		go workers(i, &wg)
	}

	wg.Wait() // Wait until all workers call Done()

	fmt.Println("All workers finished")
}
