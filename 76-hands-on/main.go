package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(251)

	fmt.Printf("The name of the var is x and the value is %v\n", x)
	if x >= 0 && x <= 100 {
		fmt.Println("between 0 and 100")
	}
	if x > 100 && x <= 200 {
		fmt.Println("between 101 and 200")
	}
	if x > 200 && x <= 250 {
		fmt.Println("between 200 and 250")
	}
}
