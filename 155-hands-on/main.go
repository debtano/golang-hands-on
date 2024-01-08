package main

import "fmt"

func main() {
	defer f1()
	defer f2()
	defer f3()
	f4()
}

func f1() {
	fmt.Println("i am f1")
}

func f2() {
	fmt.Println("i am f2")
}

func f3() {
	fmt.Println("i am f3")
}

func f4() {
	fmt.Println("i am f4")
}
