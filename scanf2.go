package main

import "fmt"

func main() {
	var name string
	var num int
	var num_float float32
	var val_bool bool

	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &num)
	fmt.Scanf("%g", &num_float)
	fmt.Scanf("%t", &val_bool)

	fmt.Printf("%s %d %g %t\n", name, num, num_float, val_bool)
}
