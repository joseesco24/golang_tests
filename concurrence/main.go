package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	// Se usa defer para indicar al wait group que todo esta listo una vez finalice la funcion.

	defer wg.Done()

	fmt.Println(text)

}

func main() {

	// Configuracion del wait group.

	var wg sync.WaitGroup

	fmt.Println("Hola")

	// Se agrega una go routine al wait group.

	wg.Add(1)

	// Se ejecuta la funcion say y se envia el texto junto con el puntero del wait group.

	go say("Mundo", &wg)

	// Se indica a la funcion principal que espere hasta que todas las tareas del wait group finalicen.

	wg.Wait()

}
