package main

import "fmt"

func main() {

	// Ejemplo de defer.

	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break.

	for i := 0; i < 10; i++ {

		fmt.Println(i)

		// Continue.

		if i == 4 {
			fmt.Println("i es 4")
			continue
		}

		// Break.

		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

}
