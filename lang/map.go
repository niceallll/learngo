package main

import (
	"fmt"
	"strings"
)

func lengthOfNoRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	s := "xiaoma1123123"
	fmt.Println(strings.Split(s, " "))
	fmt.Println(lengthOfNoRepeatingSubStr("我是下吗小马小"))

}
