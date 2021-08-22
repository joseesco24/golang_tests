package main

import "fmt"

func main() {

	// Declarando el map.

	mapa := make(map[string]int)

	// Agregra valores al map.

	mapa["pepito"] = 20
	mapa["jose"] = 14

	// Imprimiendo el map.

	fmt.Println(mapa)

	// Recorriendo el map con range.

	for llave, valor := range mapa {
		fmt.Println(llave, valor)
	}

	// Econtrar un valor del map y confirmas si la llave existe.

	value, ok := mapa["jose"]
	fmt.Println(value, ok)

}
