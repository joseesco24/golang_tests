package main

import "fmt"

// Intanciamiento de la inerfaz.

type figuras2D interface {
	area() float32
}

// Intanciamiento de los structs.

type cuadrado struct {
	base float32
}

type rectangulo struct {
	base   float32
	altura float32
}

// Intanciamiento de los metodos compartidos.

func (c cuadrado) area() float32 {
	return c.base * c.base
}

func (r rectangulo) area() float32 {
	return r.base * r.altura
}

// Intanciamiento de la funcion que usa la interfaz para usar los metodos comprtidos.

func calculate(f figuras2D) {
	fmt.Println("Area ", f.area())
}

func main() {

	// Intanciamiento de los structs.

	var myCuadrado = cuadrado{base: 4}
	var myRectnagulo = rectangulo{base: 4, altura: 2}

	// Uso de la fucncion con metodos copartidos.

	calculate(myCuadrado)
	calculate(myRectnagulo)

}
