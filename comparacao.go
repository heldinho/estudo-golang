package main

import "fmt"

func main() {
	var a, b int
	a = 2
	b = 3
	fmt.Println("a > b", a > b)
	fmt.Println("a >= b", a >= b)
	fmt.Println("a < b", a < b)
	fmt.Println("a <= b", a <= b)
	fmt.Println("a == b", a == b)
	fmt.Println("a != b", a != b)

	for a <= b {
		fmt.Println("a", a)
		a++
	}
}
