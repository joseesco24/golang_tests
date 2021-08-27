package main

import "fmt"

// Instanciamiento del struct.

type car struct {
	brand string
	year  int
}

func newCar(brand string, year int) *car {
	return &car{brand: brand, year: year}
}

func main() {

	// Metodo 4 (simulacion de contructor)

	car4 := newCar("Mazda", 2024)

	fmt.Println(*car4)

}
