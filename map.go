package main

import (
	"fmt"
	"sort"
)

func main() {
	bulan := map[int]string{
		1:  "Januari",
		2:  "Februari",
		3:  "Maret",
		4:  "April",
		5:  "Mei",
		6:  "Juni",
		7:  "Juli",
		8:  "Agustus",
		9:  "September",
		10: "Oktober",
		11: "November",
		12: "Desember",
	}

	//var keys []int
	//for k := range bulan{
	//	keys = append(keys, k)
	//}
	//sort.Ints(keys)
	keys := sortedKeysAsc(bulan)
	for _, k := range keys {
		fmt.Println("Bulan ke ", k, "adalah", bulan[k])
	}
	fmt.Println("-------------")
	fmt.Println("Reverse Order Map")
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		fmt.Println("Bulan ke ", k, "adalah", bulan[k])
	}
	fmt.Println("-------------")
	fmt.Println("Default Order Map (asc)")
	for _, k := range sortedKeysAsc(bulan) {
		fmt.Println("Bulan ke ", k, "adalah", bulan[k])
	}
	fmt.Println("-------------")
	var value, isExist = bulan[12]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Not Exist")
	}

	fmt.Println("------------------")
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	fmt.Println("----------------")
	var data = []map[string]string{
		{"nama": "Ryan", "gender": "male", "skin": "yellow"},
		{"address": "Jakarta", "id": "201"},
		{"hobbies": "movie"},
	}

	sliceKeys := make([]string, cap(data))

	i := 0
	for k := range data {
		//sliceKeys[i] = k
		//fmt.Println(sliceKeys[i])
		fmt.Println(data[i])
		for index := range data[k] {
			sliceKeys = append(sliceKeys, data[k][index])
		}
		i++
	}

	//fmt.Println(data[0]["nama"])
}

type reversedKeys struct {
	Index int
	Value string
}

func sortedKeysAsc(m map[int]string) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

func mapkey(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
