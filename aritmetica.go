package main

import "fmt"

func main() {
	var a int = 23
	var b int = 7

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	if 7 > 10 {
		fmt.Println("é maior")
	} else if 7 == 10 {
		fmt.Println("é igual")
	} else {
		fmt.Println("não é maior")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for {
		if i > 10 {
			break
		}
		i++
	}

	fmt.Println(i)
}
