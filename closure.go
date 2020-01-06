package main

import "fmt"

func main() {
	getMinMax := func(n []int) (int, int) {
		var min, max int
		for index, number := range n {
			switch {
			case index == 0:
				max, min = number, number
			case number > max:
				max = number
			case number < min:
				min = number
			}
		}
		return min, max
	}

	numbers := []int{100, 90, 4, 5, 43, 543, 5, 82, 2910, 3930, 99, 3234}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nMin : %v\nMax : %v", numbers, min, max)
}
