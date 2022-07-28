package main

import "fmt"

func main() {
	tabuleiro := [][]string{
		[]string{"x", "-", "-"},
		[]string{"o", "x", "o"},
		[]string{"-", "o", "x"},
	}

	for i := 0; i < 3; i++ {
		fmt.Println(tabuleiro[i])
	}
}
