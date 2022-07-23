package main

import "fmt"

func main() {
	var a, b int = 1, 2
	x := true
	y := false
	fmt.Println(x && y)
	fmt.Println(x || y)
	fmt.Println(!(x && y))
	fmt.Println(a > b && b > a)
	fmt.Println(b > a && b > a)
	fmt.Println(a > b || b > a)
	fmt.Println(b > a || b > a)
	fmt.Println(!(a > b && b > a))
	fmt.Println(!(a > b || b > a))
}
