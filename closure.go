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
	fmt.Printf("data : %v\nMin : %v\nMax : %v\n", numbers, min, max)

	immediatelyInvokedFunctionExpress()

	total, result := findmax(numbers, 100)
	fmt.Printf("total numbers : %d\nResult : %v", total, result())
}

func immediatelyInvokedFunctionExpress() {
	var numbers = []int{2, 3, 1, 2, 4, 2, 5, 6, 4, 46, 8, 98}

	var newNumber = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(98)
	fmt.Printf("Original Number : %v\n", numbers)
	fmt.Printf("Filter Number : %v\n", newNumber)
}

func findmax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}
