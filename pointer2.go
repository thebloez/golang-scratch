package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")

	numberA = 100

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("---Pointer Sebagai Parameter---")
	number := 4
	fmt.Println("before : ", number)

	changePointerAsParameter(&number, 10)
	fmt.Println("after : ", number)
}

func changePointerAsParameter(original *int, val int) {
	*original = val
}
