package main

import "fmt"

func main() {

	maps := make(map[string]int)
	maps["Kamran"] = 10
	maps["Ali"] = 20
	maps["Ahmed"] = 30
	fmt.Println("all amps",maps)
	delete(maps, "Kamran")

	fmt.Println("Marks of Kamran", maps["Ali"])


	for index, value := range maps {
		fmt.Println("Index is", index)
		fmt.Println("Value is", value)

	}

}
