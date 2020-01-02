package main

import (
	"fmt"
)

type Dog struct {
	Animal
}
type Animal struct {
	Age int
}

func (a *Animal) Move() {
	fmt.Println("Animal moved")
}
func (a *Animal) SayAge() {
	fmt.Printf("Animal age: %d\n", a.Age)
}
func main() {
	d := Dog{}
	d.Age = 7
	d.Move()
	d.SayAge()
}
