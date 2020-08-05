package main

import "fmt"

func main() {
	numbers := []int{2,3,1,12,3,4,3,0,2,3,12,4,32,3,4,3,234,32}

	maxFiltered, filteredNumber := func(min int) ([]int, []int) {
		var r []int
		var n []int
		for _, e := range numbers {
			if e <= min {
				n = append(n, e)
				continue
			}

			r = append(r,e)
		}
		return r, n
	}(4)
	fmt.Println("original number : ", numbers)
	fmt.Println("maxFiltered number :", maxFiltered)
	fmt.Println("Filtered Number :", filteredNumber)

	length, max := findMax(numbers, 4)
	fmt.Printf("length number : %d\nMax Number : %v", length, max())
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _,e := range numbers{
		if e >= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}
