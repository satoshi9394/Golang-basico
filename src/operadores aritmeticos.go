package main

import (
	"fmt"
	"math"
)

func main() {

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50
	// suma
	result := x + y
	fmt.Println("Suma:", result)

	// resta
	result = y - x
	fmt.Println("Resta:", result)

	// multiplicacion
	result = y * x
	fmt.Println("Multiplicacion:", result)

	// division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Retos calcula el area de:
	// -Rectangulo, Trapecio y de un circulo

	areaRectangulo := x * y
	fmt.Println("Area de un Rectangulo:", areaRectangulo)

	//Trapecio
	b := 10
	h := 15
	B := 12
	areaTrapecio := ((B + b) * h) / 2
	fmt.Println("Area de un Trapecio:", areaTrapecio)

	// Circulo
	radio := 10
	pi := 3.1416
	areaCirculo := pi * math.Pow(float64(radio), 2)
	println("Area de un Circulo:", areaCirculo)
}
