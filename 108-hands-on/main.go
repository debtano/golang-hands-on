package main

import "fmt"

func main() {

	ia := []int{40, 41, 42, 43, 44}
	for _, v := range ia {
		fmt.Printf("value is %v and type is %T\n", v, v)
	}
}
