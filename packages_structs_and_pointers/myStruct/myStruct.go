package myStruct

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc Pc) Ping() {
	fmt.Println(myPc.Brand, "pong")
}

func (myPc *Pc) DuplicateRam() {
	myPc.Ram *= 2
}
