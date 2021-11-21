package main

import "fmt"

func normalFuncion(messages string) {
	fmt.Println(messages)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFuncion("hello word")
	tripeArgument(1, 2, "hola")
	value := returnValue(2)
	fmt.Println("Value:", value)
	value1, value2 := doubleReturn(2)
	fmt.Printf("El valor 1 es: %d y el valor 2 es: %d\n", value1, value2)
	valueUno, _ := doubleReturn(2)
	fmt.Println("valor es:", valueUno)
}
