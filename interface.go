package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Phi * math.Pow(l.jariJari(), 2)
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("==== Persegi ====")
	fmt.Println("luas       :", bangunDatar.luas())
	fmt.Println("keliling   :", bangunDatar.keliling())
	fmt.Println("==== ======= ====")

	bangunDatar = lingkaran{14.0}
	var bangunLingkaran lingkaran = bangunDatar.(lingkaran)
	fmt.Println("==== Lingkaran ====")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunLingkaran.jariJari())
	fmt.Println("==== ======= ====")

}
