package main

import (
	"fmt"
	"log"
	"strconv"
)

func condicionalesGo() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// with and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// with or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad el or")
	}

	// convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	value3 := 5
	value4 := 6
	if isParNumber(value3) {
		fmt.Println("Es el value3 un par")
	} else {
		fmt.Println("No es el value3 un par")
	}

	if isParNumber(value4) {
		fmt.Println("Es el value4 un par")
	} else {
		fmt.Println("No es el value4 un par")
	}

	user1 := "Alpha"
	pass1 := "Mi_password_2564"

	user2 := "Beta"
	pass2 := "Mi_password_256"

	if isValidUser(user1, pass1) {
		fmt.Println("El usuario user1 es correcto")
	} else {
		fmt.Println("El usuario user1 es incorrecto")
	}

	if isValidUser(user2, pass2) {
		fmt.Println("El usuario user2 es correcto")
	} else {
		fmt.Println("El usuario user2 es incorrecto")
	}

}

func isParNumber(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "Mi_password_2564"
}
