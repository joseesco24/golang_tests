package main

import "fmt"

func normal_function(message string) {
	fmt.Println(message)
}

func tripla_argument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func simple_return_function(a int) int {
	return a * 2
}

func double_return_function(a int) (c, d int) {
	return a, a * 2
}

func main() {

	value_1, value_2 := double_return_function(4)
	fmt.Println(value_1, value_2)

	value_4, _ := double_return_function(4)
	fmt.Println(value_4)

	fmt.Println(double_return_function(4))

	value_3 := simple_return_function(4)
	fmt.Println(value_3)

	fmt.Println(simple_return_function(4))

	normal_function("Hola Mundo")

	tripla_argument(1, 2, "Hola")

}
