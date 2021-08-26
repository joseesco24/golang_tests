package main

import (
	"fmt"
	pk "main/myStruct"
)

func main() {

	// Se instancia el struct.

	var myPc = pk.Pc{Ram: 16, Disk: 240, Brand: "MSI"}
	fmt.Println(myPc)

	// Se llama la funcion ping.

	myPc.Ping()

	// Se llama a la funcion duplicate dos veces.

	myPc.DuplicateRam()
	fmt.Println(myPc.Ram)

	myPc.DuplicateRam()
	fmt.Println(myPc.Ram)

}
