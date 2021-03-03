package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 66.06
	var reflectVal = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectVal.Type())

	if reflectVal.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectVal.Int())
	}

	reflectStr := "test reflect"

	fmt.Println("tipe :", reflectStr)
}
