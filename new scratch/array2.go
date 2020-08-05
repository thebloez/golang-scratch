package main

import (
	"fmt"
)

func main() {
	var a [4]string

	a[0] = "Ryan"
	a[1] = "Safary"
	a[2] = "Hidayat"

	b := [...]int{1, 2, 3}

	fmt.Println("b length : ", len(b))
	fmt.Println("b content is : ", b)

	// array multidimensi
	var numbers1 = [2][3]int{{3, 2, 3}, {1, 2, 3}}
	fmt.Println(numbers1)

	fmt.Println("-------------------")

	arrNum := [...]int{3,4,5,432,5,6,0,7}

	fmt.Println("From Looping")
	for i := 0; i < len(arrNum); i++ {
		fmt.Println("index ke : ",i,", isinya : ", arrNum[i])
	}
	fmt.Println("length : ", len(arrNum))

	// for each array
	for i,a := range arrNum {
		fmt.Println("index ke : ", i ,", isinya : ", a)
	}
}
