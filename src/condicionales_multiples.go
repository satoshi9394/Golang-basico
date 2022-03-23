package main

import (
	"fmt"
)

func main() {
	// modulo := 6 % 2

	switch modulo := 6 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es inpar")
	}

	// sin condiciones
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}

}
