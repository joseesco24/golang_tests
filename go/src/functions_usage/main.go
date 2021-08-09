package main

import "fmt"

func regular_function(message string) {
	fmt.Println(message)
}

func function_with_three_returns(a, b int, c string) {
	fmt.Println(a, b, c)
}

func regular_function_with_return(a int) int {
	return a * 2
}

func function_with_two_returns(a int) (c, d int) {
	return a, a * 2
}

func main() {

	value_1, value_2 := function_with_two_returns(4)
	fmt.Println(value_1, value_2)

	value_4, _ := function_with_two_returns(4)
	fmt.Println(value_4)

	fmt.Println(function_with_two_returns(4))

	value_3 := regular_function_with_return(4)
	fmt.Println(value_3)

	fmt.Println(regular_function_with_return(4))

	regular_function("Hola Mundo")

	function_with_three_returns(1, 2, "Hola")

}
