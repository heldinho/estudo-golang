package main

import "fmt"

func main() {
	tipo1()
	tipo2()
	tipo3()
}

func tipo1() {
	var idade, altura int = 32, 185
	var nome string = "Tipo1"
	fmt.Println(nome, idade, altura)
}

func tipo2() {
	var (
		idade  int    = 32
		altura int    = 185
		nome   string = "Tipo2"
	)
	fmt.Println(nome, idade, altura)
}

func tipo3() {
	idade := 32
	altura := 185
	nome := "Tipo3"
	fmt.Println(nome, idade, altura)
}
