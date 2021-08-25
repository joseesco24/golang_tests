package main

import "fmt"

func main() {

	// Ejercicio.

	// Calcular el áera del cuadrado.

	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	// Declaracion de variables.

	x := 10
	y := 50

	// Operaciones aritmeticas.

	// Suma.
	result := x + y
	fmt.Println("Suma:", result)

	// Resta.
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicación.
	result = y * x
	fmt.Println("Multiplicación:", result)

	// División.
	result = y / x
	fmt.Println("División:", result)

	// Modulo.
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental.
	x++
	fmt.Println("Incremental:", x)

	// Decremental.
	x--
	fmt.Println("Incremental:", x)

}
