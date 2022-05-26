package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	//contes, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contes)
	//}
	if contes, err := ioutil.ReadFile(filename); err == nil {
		fmt.Println(string(contes))
	} else {
		fmt.Println(err)
	}

	fmt.Println(
		grade(0),
		grade(1),
		grade(91),
		grade(59),
		grade(100),
		grade(-1),
	)
}
