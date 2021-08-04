package main

import "fmt"

func main() {

	// Rectangulo

	const base_rectangulo float32 = 14.1252
	const altura_rectangulo float32 = 18

	area_rectangulo := base_rectangulo * altura_rectangulo

	fmt.Println("El area del rectangulo es de:", area_rectangulo)

	// Trapecio

	const base_mayor_trapecio float32 = 14
	const base_menor_trapecio float32 = 18
	const altura_trapecio float32 = 10

	area_trapecio := ((base_mayor_trapecio + base_menor_trapecio) / 2) * altura_trapecio

	fmt.Println("El area del trapecio es de:", area_trapecio)

	// Circulo

	const pi float32 = 3.141642
	const radio_circulo float32 = 9

	area_circulo := pi * (radio_circulo * radio_circulo)

	fmt.Println("El area del circulo es de:", area_circulo)

}
