package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for z := 0; z < 100; z++ {

		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("value of x is %v\t value of y is %v\n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("both are less than 4")
		case x > 6 && y > 6:
			fmt.Println("both are greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater or equal to 4 and less or equal than 6")
		case y != 5:
			fmt.Println("y is different than 5")
		default:
			fmt.Println("none of the cases were meet")
		}
	}
}
