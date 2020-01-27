package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"Ryan", "Safary", "Hidayat"}
	var dataContainsF = filter(data, func(each string) bool {
		return strings.Contains(each, "f")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) > 4
	})

	fmt.Println("data asli \t\t\t\t:", data)
	fmt.Println("filter ada huruf \"f\"\t:", dataContainsF)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
