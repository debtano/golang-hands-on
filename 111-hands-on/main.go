package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xa := append(x, 52)
	fmt.Println(xa)

	xb := append(xa, 53, 54, 55)
	fmt.Println(xb)

	y := []int{56, 57, 58, 59, 60}
	xc := append(xb, y...)
	fmt.Println(xc)
}
