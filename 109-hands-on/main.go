package main

import "fmt"

func main() {

	ai := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a := ai[:5]
	fmt.Println(a)

	b := ai[5:]
	fmt.Println(b)
	c := ai[2:7]
	fmt.Println(c)
	d := ai[1:6]
	fmt.Println(d)

	for i, v := range ai {
		fmt.Printf("at position %v value is %v of type %T\n", i, v, v)
	}

}
