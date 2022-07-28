package main

import (
	"fmt"
)

func main() {
	var nomes = [3]string{
		"Ana",
		"Jose",
		"Maria",
	}
	fmt.Printf("len=%d, cap=%d\n", len(nomes), cap(nomes))
	fmt.Println(nomes)
}
