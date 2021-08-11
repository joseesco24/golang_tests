package main

import "fmt"

func main() {

	// Declaración de un array.

	var array [4]int

	// Cambio de los valores del array.

	array[0] = 1
	array[1] = 2

	// Impresión de los valores del array.

	fmt.Println(array)

	// Impresión de las especificaciones del array.

	// Número de elementos en el array.

	fmt.Println(len(array))

	// Capacidad máxima del array.

	fmt.Println(cap(array))

	// Declaración de un slice.

	slice := []int{0, 1, 2}

	// Agregar elementos al slice.

	slice = append(slice, 3, 4, 5)

	// Impresión de los valores del slice.

	fmt.Println(slice)

	// Número de elementos en el slice.

	fmt.Println(len(slice))

	// Capacidad máxima del slice.

	fmt.Println(cap(slice))

	// Métodos de un slice.

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append de un nuevo slice.

	new_slice := []int{6, 7, 8}
	slice = append(slice, new_slice...)

	fmt.Println(slice)

}
