package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}

func main() {

	c := make(chan string, 2)

	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	// Range y close

	close(c)

	// c <- "mensaje 3"

	for message := range c {
		fmt.Println(message)
	}

	// select

	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje email 1", email1)
	go message("mensaje email 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)

		}
	}

}
