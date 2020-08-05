package main

import (
	"fmt"
	"math"
)

func main() {
	avg, index := calculate(2,3,4,5,7,6,5,5,3)
	msg := fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
	fmt.Println(index)
	area, circum := calculating(2,3,4.2,4)
	msg2 := fmt.Sprintf("Calculating area : %.2f", area)
	msg3 := fmt.Sprintf("Calculating area : %.2f", circum)
	fmt.Println(msg2)
	fmt.Println(msg3)

}

func calculate(numbers ...int) (float32, int) {
	total := 0
	index := 0
	for idx, number := range numbers {
		total += number
		index = idx
	}
	avg := float32(total) / float32(len(numbers))
	return avg, index
}

func calculating(d ...float64) (area float64, circumference float64) {
	total := 0
	for _,number := range d {
		number+=number
		total = int(number)
	}
	area = math.Pi * math.Pow(float64(total/2), 2)
	circumference = math.Pi * float64(total)

	return
}
