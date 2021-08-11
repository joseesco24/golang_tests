package main

import (
	"fmt"
	"strings"
)

func main() {

	// Declarando el slice.

	slice := []string{"hola", "que", "hace"}

	// Recorriendo el slice usando indice y valor.

	for indice, valor := range slice {
		fmt.Println(indice, valor)
	}

	// Recorriendo el slice usando solo el valor.

	for _, valor := range slice {
		fmt.Println(valor)
	}

	// Recorriendo el slice usando solo el indice.

	for indice := range slice {
		fmt.Println(indice)
	}

	// Recorriendo un string.

	for indice, valor := range strings.ToLower("hOla") {
		fmt.Println(indice, string(valor))
	}

}
