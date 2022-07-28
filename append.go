package main

import "fmt"

func main() {
	var frutas []string
	frutas = append(frutas, "Maçã")
	frutas = append(frutas, "Laranja")
	frutas = append(frutas, "Morango", "Limão", "Uva", "Pera", "Manga", "Abacaxi")

	fmt.Println(frutas)
	fmt.Printf("%d %d\n", len(frutas), cap(frutas))
}
