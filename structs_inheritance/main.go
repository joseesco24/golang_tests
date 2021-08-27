package main

import "fmt"

// Instanciamiento del struct persona.

type Person struct {
	name     string
	birdYear int
}

// Intanciamiento del struct empleado.

type Employee struct {
	id int
}

// Intanciamiento del struct empleado de tiempo completo.

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {

	// Creacion del objeto de empleado de tiempo completo.

	newFullTimeEmployee := FullTimeEmployee{}

	// Llenando los atributos del empleado.

	newFullTimeEmployee.name = "Jose"
	newFullTimeEmployee.birdYear = 1998
	newFullTimeEmployee.id = 14

	// Imprimiendo la estructura del empleado.

	fmt.Println(newFullTimeEmployee)

}
