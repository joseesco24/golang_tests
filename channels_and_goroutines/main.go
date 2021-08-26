package main

import "fmt"

// Al usar channels en funciones se puede especificar si el canal es de entrada o salida con canal<- o <-canal, respectivamente.

func say(text string, c chan<- string) {

	// Se ingresa el dato al canal.
	c <- text

}

func main() {

	// Creacion del canal, si no se indican el numero de datos que recivira el canal recivira de forma dinamica los datos.

	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	// Extrae el dato del canal.

	fmt.Println(<-c)

}
