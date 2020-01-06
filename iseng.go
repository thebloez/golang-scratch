package main

import "fmt"

func main() {

	fmt.Println("Hello Go-Lang")
	fmt.Println(average(1, 2, 3, 4, 5, 6))
}

func average(numbers ...int) (avg float64) {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	avg = float64(total) / float64(len(numbers))
	return
}
