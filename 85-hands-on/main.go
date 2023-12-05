package main

import "fmt"

func main() {
	// x := 1
	for x := 0; x <= 100; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Printf("even number : %v\n", x)
	}
}
