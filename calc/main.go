package main

import "fmt"

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(sub(3, 1))
	fmt.Println(times(2, 3))
	fmt.Println(divide(6, 2))
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}
