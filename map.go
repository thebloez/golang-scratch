package main

import (
	"fmt"
	"sort"
)

func main() {
	bulan := map[int]string{
		1 : "Januari",
		2 : "Februari",
		3 : "Maret",
		4 : "April",
		5 : "Mei",
		6 : "Juni",
		7 : "Juli",
		8 : "Agustus",
		9 : "September",
		10 : "Oktober",
		11 : "November",
		12 : "Desember",
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
	for _, k := range sortedKeysAsc(bulan){
		fmt.Println("Bulan ke ", k, "adalah", bulan[k])
	}
	fmt.Println("-------------")
	var value, isExist = bulan[12]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Not Exist")
	}

	fmt.Println()
}

type reversedKeys struct {
	Index int
	Value string
}

func sortedKeysAsc(m map[int]string) ([]int) {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}


