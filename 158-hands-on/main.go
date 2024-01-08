package main

import "fmt"

func Add(a int, b int) int {
	return a * b
}

func main() {
	t := Add(5, 5)
	fmt.Println(t)
}
