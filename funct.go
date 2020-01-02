package main

import "fmt"

func main() {
	var (
		rs string
		is int
	)
	rs, is = multipleReturn(2,3)
	fmt.Println(rs, is)
}

func multipleReturn(i1 int, i2 int) (string, int)  {
	returnOfString := "hasil penjumlahan adalah "
	returnOfInt := i1 + i2
	return returnOfString, returnOfInt
}
