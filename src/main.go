package main

import (
	"fmt"

	myPk "Golang-basic/src/mypackage"
)

func main() {
	var myCar myPk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2011
	fmt.Println(myCar)

	myPk.PrintMsg("hola platzi")
}
