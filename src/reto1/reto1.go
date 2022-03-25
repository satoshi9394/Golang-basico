package reto1

import "fmt"

// mi pc strutur
type pc struct {
	ram   int
	disk  int
	brand string
}

// Ping para ver la marga y un mensaje de pong
func (myPc pc) Ping() {
	fmt.Println(myPc.brand, "Pong")
}

// DuplicateRam para duplicar ram
func (myPc *pc) DuplicateRam() {
	myPc.ram = myPc.ram * 2
}

// New crear un nueva pc
func New(rm int, dsk int, brn string) pc {
	myPc := pc{ram: rm, disk: dsk, brand: brn}
	return myPc
}

// para imprimir mi pc
func (myPc pc) String() string {
	return fmt.Sprintf("tengo %d Gb Ram, %d GB de disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}
