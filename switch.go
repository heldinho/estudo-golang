package main

import "fmt"

func main() {
	var nome string
	var inicio int = 1

	for inicio == 1 {
		fmt.Print("Digite um nome: ")
		fmt.Scanf("%s", &nome)

		switch nome {
		case "ana":
			fmt.Println("\n=============================\n=========> é a ana <=========\n=============================\n")
		case "joao":
			fmt.Println("\n==============================\n=========> é a joao <=========\n==============================\n")
		default:
			fmt.Println("\n=================================\n=========> não conheço <=========\n=================================\n")
		}

		fmt.Print("Testar de novo\n  1 - (Sim)\n  0 - (Não)\nDigite: ")
		fmt.Scanf("%d", &inicio)
	}

	fmt.Println("fim do programa")
}
