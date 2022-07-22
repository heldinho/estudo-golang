package main

import "fmt"

func soma(x, y int) int {
	return x + y
}

func calcular(a int, nome string) (int, int, string) {
	var quadrado int = a * a
	var cubo int = a * a * a

	return quadrado, cubo, nome
}

func main() {
	fmt.Println(calcular(2, "Helder Passos"))
}
