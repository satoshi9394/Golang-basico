package main

import (
	"fmt"
	"strings"
)

func isPalindomo(text string) {
	var textRevers string
	for i := len(text) - 1; i >= 0; i-- {
		textRevers += string(text[i])
	}

	if text == textRevers {
		fmt.Println("es palindromo")
	} else {
		fmt.Println("no es palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindomo("ama")
	isPalindomo("amor a roma")
	isPalindomo("tu mama")
	var palabra string
	fmt.Scan(&palabra)
	minus := strings.ToLower(palabra)
	isPalindomo(minus)
}
