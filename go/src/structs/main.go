package main

import "fmt"

// Instanciamiento del struct.

type car struct {
	brand string
	year  int
}

func main() {

	// Creacion del Struct.

	my_car := car{brand: "hyundai", year: 2020}

	// Impresion del Struct.

	fmt.Println(my_car)

	// Creacion del struct.

	var other_car car
	other_car.brand = "Ferrari"

	// Impresion del Struct.

	fmt.Println(other_car)

}
