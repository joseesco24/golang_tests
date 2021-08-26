package main

import "fmt"

func main() {

	// Intanciamiento de la lista de interfaces.

	myInterface := []interface{}{"Hola", 12, 4.9}

	// gregando un nuevo elemento al slice.

	myInterface = append(myInterface, "holi")

	fmt.Println(myInterface...)

}
