package main

import "fmt"

func main() {
	fmt.Println("------------------")
	fmt.Println("Closure")
	closure()

	fmt.Println("------------------")
	fmt.Println("Immediately-Invoked Function Expression (IIFE)")
	immediatelyInvokedFunctionExpress()

	fmt.Println("------------------")
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findmax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

func closure() {
	// closure
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{75, 54, 4, 7, 7, 3, 55, 2, 6}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nMin : %v\nMax : %v\n", numbers, min, max)
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
	}(8)
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
