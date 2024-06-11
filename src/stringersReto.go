package main

import (
	rt "Golang-basic/src/reto1"
	"fmt"
)

func stringersReto() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)
	myPc := rt.New(16, 200, "msi")

	fmt.Println(myPc.String())
	myPc.Ping()
	fmt.Println(myPc.String())
	myPc.DuplicateRam()
	fmt.Println(myPc.String())
	myPc.DuplicateRam()
	fmt.Println(myPc.String())
}
