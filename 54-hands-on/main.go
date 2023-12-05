package main

import (
	"fmt"
)

var outside string = "outsider"

const out string = "out constant"

func main() {
	a := "inside"
	b := 0.0

	fmt.Println(outside, "\t", out, "\t", a, "\t", b)

	// fmt.Println(puppy.)
	fmt.Println()
}
