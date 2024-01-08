package main

import "fmt"

func main() {

	si := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(foo(si...))
	fmt.Println(bar(si))

}

func foo(si ...int) int {
	sum := 0
	for _, v := range si {
		sum += v
	}

	return sum

}

func bar(osi []int) int {
	sum := 0
	for _, v := range osi {
		sum += v
	}

	return sum
}
