package main

import "fmt"

func main() {

	// Declaracion de variables.

	valor_1 := 5 % 2

	// Switch clasico.

	switch valor_1 {
	case 0:
		fmt.Println("Si es par")
	default:
		fmt.Println("No es par")
	}

	// Switch declarado en la misma a linea.

	switch valor_2 := 4 % 2; valor_2 {
	case 0:
		fmt.Println("Si es par")
	default:
		fmt.Println("No es par")
	}

	// Switch condicional.

	valor_3 := 100

	switch {
	case valor_3 > 100:
		fmt.Println("Mayor")
	case valor_3 < 100:
		fmt.Println("Menor")
	default:
		fmt.Println("Igual")
	}

}
