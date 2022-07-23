package main

import "fmt"

type helder int

var num helder = 0

func main() {
	fmt.Printf("%T, %v", num, num)
}
