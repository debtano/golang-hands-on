package main

import "fmt"

func main() {
	for x := 1; x <= 5; x++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("inner loop is : %v\t\t outer loop is : %v\n", x, j)
		}
	}
}
