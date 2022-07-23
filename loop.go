package main

import "fmt"

func loopQuadrado() {
	var quadrado int = 2
	for quadrado <= 1024 {
		quadrado *= quadrado
	}
	fmt.Println(quadrado)
}

func loopWhile() {
	var soma int = 2
	for soma < 600 {
		soma += soma
	}
	fmt.Println(soma)
}

func loopFor() {
	var soma int = 0
	var i int
	for i = 0; i < 10; i++ {
		soma += i
	}
	fmt.Println(soma)
}

func main() {
	loopFor()
	loopWhile()
	loopQuadrado()
}
