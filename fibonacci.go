package main

import "fmt"

func fibonacci(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func main() {
	var i int
	n := 45
	fmt.Println("Digite a quantidade de termos da sequência de Fibonacci:")
	fmt.Println("A sequência de Fibonacci e:")
	for i = 0; i < n; i++ {
		fmt.Println(fibonacci(i + 1))
	}
}
