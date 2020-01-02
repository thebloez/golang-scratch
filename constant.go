package main

import "fmt"

type tipeString string
type tipeInt int

const myUntypedConstString  = "myUntypedConstString"
const myUnutypedConstInteger = 28
const c = 1 << 5

func main() {
	hello := "hello"
	var tips tipeString = tipeString(hello)
	var myInt int = 9
	var types tipeInt = tipeInt(myInt)

	var myStr = myUntypedConstString
	tambahinConst := myUnutypedConstInteger + 2

	fmt.Println(tips)
	fmt.Printf("%d\n",types)
	fmt.Println(myStr)
	fmt.Printf("%d\n", tambahinConst)
	fmt.Printf("%d\n", c)


	var result = 25/2.0
	fmt.Printf("result is %f which is of type %T\n", result, result)

	var sabaraha = 4.5 + (10 - 5) * (3 + 2)/2 // 4.5 + (5) * (5) / 2 | 4.5 + 25 / 2 | 4.5 + 12
	fmt.Printf("Result = %f\n", sabaraha)
	//fmt.Printf("%f\n",)
	fmt.Printf("%f\n",4.5 + 12)

	//-------------

	var resultCorrect = 4.5 + (10 - 5) * (3 + 2)/2.0 // 4.5 + (5) * (5) / 2.0 | 4.5 + 25 / 2.0
	//var result2 = 4.5 + float64((10 - 5) * (3 + 2))/2 // 4.5 + 25.0 / 2
	fmt.Printf("%f\n", resultCorrect)
	fmt.Printf("%f\n", 4.5+25.0/2)
	fmt.Printf("%f\n",29.5 / 2)

}

