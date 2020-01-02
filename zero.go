package main

import "fmt"

func main() {
	var (
		firstname, lastname string
		umur int
		tinggi float64
		isActive bool
	)

	fmt.Printf("firstname : %s, lastname : %s, umur : %d, tinggi : %f, aktif : %t",
		firstname,lastname, umur, tinggi,isActive)

	basicOperator()
}

func basicOperator() {
	myVarInt := 20
	var myUint uint = 13
	myHexNumber := 0xdd
	myOctal := 034
	myFloat := 0.6413
	fmt.Printf("\n")
	fmt.Printf("myVarInt : %#x, myUINT : %d, myHexNumber = %#x, myOctal = %#o\n", myVarInt, myUint,
		myHexNumber, myOctal)

	fmt.Printf("myFloat : %f, with limit : %.2f\n, int : %d", myFloat, myFloat, myVarInt)

}


