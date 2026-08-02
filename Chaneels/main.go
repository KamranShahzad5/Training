package main

import "fmt"

func worker(tasks chan string, results chan string) {
	// Receive task from main
	task := <-tasks
	fmt.Println("Worker received:", task)


	// Send result back to main
	results <- task
}

func main() {
	// Create channels
	tasks := make(chan string)
	results := make(chan string)

	// Start worker
	go worker(tasks, results)

	// Send task to worker
	tasks <- "Download File"

	// Receive result from worker
	msg := <-results

	fmt.Println("Main received:", msg)
}