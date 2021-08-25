package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Hola")
	}()

	go func(text string) {
		fmt.Println(text)
	}("Mundo")

	time.Sleep(time.Second * 2)

}
