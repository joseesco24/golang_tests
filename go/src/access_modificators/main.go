package main

import (
	pk "access_modificators/my_package"
	"fmt"
)

func main() {

	var my_car pk.Car

	my_car.Brand = "Hyundai"
	my_car.Year = 2018

	fmt.Println(my_car)
}
