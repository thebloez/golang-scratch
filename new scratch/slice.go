package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// fungsi cap
	fmt.Println("fungsi cap")
	var cFruits = fruits[0:3]
	fmt.Println(len(cFruits)) // len: 3
	fmt.Println(cap(cFruits)) // cap: 4

	var dFruits = fruits[1:4]
	fmt.Println(len(dFruits)) // len: 3
	fmt.Println(cap(dFruits)) // cap: 3
}
