package main

import "fmt"

type Creature struct {
	Name string
	Type string
}

func main() {
	c := Creature{
		Name: "Sammy",
		Type: "Shark",
	}
	fmt.Printf("%v the %v", c.Name, c.Type)
	fmt.Scanf("%d")
}
