package main

import "fmt"

func main() {
	helloMessage := "Hello "
	wordMsg := "word"
	fmt.Println(helloMessage + wordMsg)

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipos de datos
	fmt.Printf("Hello messages: %T\n", helloMessage)
}
