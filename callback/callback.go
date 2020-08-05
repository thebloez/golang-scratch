package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"wick","donald", "bebek"}

	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var filterByDigit = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data Asli \t\t\t\t: ", data)
	fmt.Println("Filter ada huruf \"o\"\t: ", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t: ", filterByDigit)
}

func filter(data []string, callback func(string) bool) []string{
	var result []string
	for _, one := range data {
		if filtered := callback(one); filtered {
			result = append(result, one)
		}
	}

	return result
}
