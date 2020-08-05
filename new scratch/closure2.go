package main

import "fmt"

func main() {
	getMinMax := func(n []int) (int, int) {
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
	var numbers = []int{3123,1231,1123,4324,4353,12342,2342,12,2143,5436,7568}
	min, max := getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v", numbers, min, max)
}
