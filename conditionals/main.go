package main

import "fmt"

func main() {

	// Declaracion de variables.

	valor_1 := 1
	valor_2 := 2

	// If clasico.

	if valor_1 == 0 {
		fmt.Println("Si es 1")
	} else if valor_2 == 2 {
		fmt.Println("Val 2 == 2")
	} else {
		fmt.Println("No es 1")
	}

	// If con and

	if valor_1 == 1 && valor_2 == 2 {
		fmt.Println("Val 1 = 1 y Val 2 = 2")
	}

	// If con or

	if valor_1 == 0 || valor_2 == 2 {
		fmt.Println("Val 1 = 0 o Val 2 = 2")
	}

}
