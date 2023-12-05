package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for y := 1; y <= 100; y++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("for iteration number %v\t\t value is 3", y)
		}
	}
}
