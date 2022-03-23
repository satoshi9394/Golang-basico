package main

import (
	"fmt"
)

func main() {
	// Defer para ejecutar antes de que muera la funcion (ultima ejecucion)
	defer fmt.Println("Hola")
	fmt.Println("mundo")

	// continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// break
		if i == 8 {
			fmt.Println("Es 8")
			break
		}
	}
}
