package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Captura del error.

	number, error := strconv.ParseInt("8", 0, 64)

	// Control del error.

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(number)
	}

}
