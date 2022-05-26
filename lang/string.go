package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "YES我爱吃饭!"
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println("RUNE count: ", utf8.RuneCountInString(s))

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}

	//冒泡排序

}
