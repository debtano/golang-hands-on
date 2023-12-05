package main

import "fmt"

func main() {

	x := 20
	for {
		fmt.Printf("value of x is %v\n", x)
		x--
		if x <= 0 {
			break
		}
	}
}
