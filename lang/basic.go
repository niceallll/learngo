package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	fmt.Println(
		cmplx.Pow(math.E, 1i*math.Pi) + 1)
}
func taring() {
	var a, b int = 1, 2
	var c int
	c = int(math.Sqrt(float64(a*b + a + b)))
	fmt.Println(c)
}
func constn() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	euler()
	taring()
	constn()
}
