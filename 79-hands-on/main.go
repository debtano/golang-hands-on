package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("value of x is %v\t value of y is %v\n", x, y)
	if x < 4 && y < 4 {
		fmt.Println("both are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("both are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater or equal to 4 and less or equal than 6")
	} else if y != 5 {
		fmt.Println("y is different than 5")
	} else {
		fmt.Println("none of the cases were meet")
	}
}
