package main

import "fmt"

func main() {

	// Declaracion de constantes.

	const pi_1 float64 = 3.14
	const pi_2 = 3.1415

	fmt.Println("Pi_1:", pi_1)
	fmt.Println("Pi_2:", pi_2)

	// Declaracion de variables.

	var altura int = 14 // Crea la variable, le asigna un tipo de dato y el valor.
	var area int        // Crea la variable y le asigna un tipo de dato sin asignar el valor.
	base := 14          // Crea la variable y se le asigna un valor sin asignar un tipo de dato.

	// Go no permite compilar o ejecutar a no ser que todas las variables instanciadas se usen.

	fmt.Println(base, altura, area)

	// Zero values, go asigna valores por defecto a las varibles cuyo valor no ha sido asignado.

	var a int
	var b string
	var c float64
	var d bool

	fmt.Println(a, b, c, d)

	// Ejercicio.

	// Calcular el áera del cuadrado.

	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

}
