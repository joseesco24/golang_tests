package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func main() {

	// Creacion del canal, si no se indican el numero de datos que recivira el canal recivira de forma dinamica los datos.

	c := make(chan string, 2)

	// Insertando datos al canal.

	c <- "Mensaje 1"
	c <- "Mensaje 2"

	// Imprimiendo datos del canal.

	fmt.Println(len(c), cap(c))

	// Close cierra el canal para que no reciva mas datos, lo adecuado es cerrar el canal solo cuando se sepa que no van a insertarse mas datos.

	close(c)

	// Recorriendo canal con range.

	for message := range c {
		fmt.Println(message)
	}

	// Ejemplo de select.

	email1 := make(chan string)
	email2 := make(chan string)

	// Insertando datos a los dos canales.

	go message("mensaje 2", email2)
	go message("mensaje 1", email1)

	// Revisando los canales.

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Recivido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Recivideo de email2", m2)
		}
	}

}
