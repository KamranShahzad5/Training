package main

import (
	"fmt"
	"os"
)

func main() {

	// 1. Create a file
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Println("File created successfully")

	// 2. Write to the file
	_, err = file.WriteString("Hello, Go!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	file.Close() // Close before reopening

	// 3. Read the file
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("\nFile Content:")
	fmt.Println(string(data))

	// 4. Append to the file
	file, err = os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	_, err = file.WriteString("This line is appended.\n")
	if err != nil {
		fmt.Println("Error appending:", err)
		return
	}
	file.Close()

}