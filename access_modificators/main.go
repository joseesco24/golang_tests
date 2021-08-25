package main

import (
	"fmt"
	pk "main/my_package"
)

func main() {

	var my_car pk.Car

	my_car.Brand = "Hyundai"
	my_car.Year = 2018

	fmt.Println(my_car)
}
