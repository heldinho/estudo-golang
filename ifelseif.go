package main

import "fmt"

func main() {
	var name string
	var num int

	fmt.Print("Digite o seu nome: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Digite um número: ")
	fmt.Scanf("%d", &num)

	if num > 10 {
		fmt.Printf("%v: %v é maior que 10\n", name, num)
	} else if num == 10 {
		fmt.Printf("%v: %v é maior que 10\n", name, num)
	} else {
		fmt.Printf("%v: %v é maior que 10\n", name, num)
	}
}
