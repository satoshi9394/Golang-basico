package mypackage

import "fmt"

type car struct {
	brand string
	year  int
}

// public en mayusculas y en minusculas privado

// CardPublic car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type cardPrivate struct {
	brand string
	year  int
}

// PrintMsg imprimir un mensaje
func PrintMsg(text string) {
	fmt.Println(text)
}

//privada
func printMsg1(text string) {
	fmt.Println(text)
}
