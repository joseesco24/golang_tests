package main

import "fmt"

func main() {

	// Ejemplo de println: agrega el salto de linea de forma automatica.

	message_1 := "message 1"
	message_2 := "message 2"

	fmt.Println(message_1, message_2)

	// Ejemplo de prinf: permite formatear el mensaje.

	mensaje_3 := "platzi"
	decimal_1 := 500

	fmt.Printf("%s tiene mas de %d cursos \n", mensaje_3, decimal_1)
	fmt.Printf("%v tiene mas de %v cursos \n", mensaje_3, decimal_1)

	// Ejemplo de sprintf: guarda el string formateado sin imprimir en consola.

	string_formateado := fmt.Sprintf("%v tiene mas de %v cursos", mensaje_3, decimal_1)
	fmt.Println(string_formateado)

	// Ejemplo de tipo de variable.
	fmt.Printf("tipo de variable de message_1: %T \n", message_1)

}
