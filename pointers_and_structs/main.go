package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "pong")
}

// Se usa el asterisco para indicar que se van a acceder los valores del pc.

func (myPc *pc) duplicateRam() {
	myPc.ram *= 2
}

func main() {

	// Se instancia el struct.

	myPc := pc{ram: 16, disk: 240, brand: "MSI"}
	fmt.Println(myPc)

	// Se llama la funcion ping.

	myPc.ping()

	// Se llama a la funcion duplicate dos veces.

	myPc.duplicateRam()
	fmt.Println(myPc.ram)

	myPc.duplicateRam()
	fmt.Println(myPc.ram)

}
