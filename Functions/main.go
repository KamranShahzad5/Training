package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}




func main() {
	ans := add(5, 6)
	fmt.Println("My answer is", ans)

	a, b := swap("hello", "world")
	fmt.Println(a, b)

}
