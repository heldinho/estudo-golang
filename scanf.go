package main

import (
	"fmt"
)

func main() {

	var name string
	var num int

	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &num)

	fmt.Printf("The word %s containing %d number of alphabets.", name, num)
}
