package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Sobre escritura del metodo String del struct.

func (myPc pc) String() string {
	return fmt.Sprintf("El pc tiene %dGB ram, %dGB disco y es un %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {

	// Se instancia el struct.

	myPc := pc{ram: 16, disk: 240, brand: "MSI"}
	fmt.Println(myPc)

}
