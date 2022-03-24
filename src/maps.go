package main

import (
	"fmt"
)

func maps() {
	// Maps similar a dicionarios en python
	m := make(map[string]int)

	m["jose"] = 14
	m["pepito"] = 20
	fmt.Println(m)

	// recorrer un map range sera aleatorio

	for i, value := range m {
		fmt.Println("llave:", i, "valor:", value)
	}

	// encontrar valor
	value := m["jose"]
	fmt.Println(value)

	// saber si una llave existe se usa el ok que devuelve el map
	value2, ok := m["josep"]
	fmt.Println(value2, ok)
}
