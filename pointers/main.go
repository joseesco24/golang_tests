package main

import "fmt"

func main() {

	// Se crea la variable a y se le asigna un valor.

	a := 50

	// Se crea la variable b y se apunta a la direccion de memoria de a.

	b := &a

	fmt.Println(a, b)

	// Se usa el asterisco para acceder al valor de la memoria en lugar de usar su direccion.

	fmt.Println(a, *b)

	// Modificando el valor de b en la memoria.

	*b = 100

	// a y b imprimen 100 ya que ambas variables apuntan a la misma direccion de memoria.

	fmt.Println(&a, &b)
	fmt.Println(*b)
	fmt.Println(a)

}
