package main

import "fmt"

func main() {
	fmt.Println(getAllSum(1, 2, 3, -4))

	// with slice
	intSlice := []int{1, 2, 3, 4, 5}
	total := getAllSum(intSlice...)
	fmt.Println(total)
}

func getAllSum(num ...int) (total int) {
	for _, elemen := range num {
		total += elemen
	}
	return
}
