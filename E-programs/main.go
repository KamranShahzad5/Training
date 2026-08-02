package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

// 1. Race Condition
func raceCondition() {
	counter = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // Race condition
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// 2. Missing Done()
func missingDone() {
	wg.Add(1)

	go func() {
		fmt.Println("Working...")
		// Missing wg.Done()
	}()

	wg.Wait() // Deadlock
}

// 3. Send Without Receiver
func sendWithoutReceiver() {
	ch := make(chan int)
	ch <- 10 // Deadlock
}

// 4. Receive Without Sender
func receiveWithoutSender() {
	ch := make(chan int)
	fmt.Println(<-ch) // Deadlock
}

// 5. Unlock Without Lock
func unlockWithoutLock() {
	mu.Unlock() // Panic
}

func main() {

	// Uncomment ONE example at a time.

	// raceCondition()

	// missingDone()

	// sendWithoutReceiver()

	// receiveWithoutSender()

	// unlockWithoutLock()
}
