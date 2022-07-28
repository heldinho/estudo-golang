package main

import "fmt"

var name string = "Helder Passos"

func main() {
	numeros := make([]int, 5)
	nome := &name
	fmt.Println(numeros)
	fmt.Println(nome)
}
