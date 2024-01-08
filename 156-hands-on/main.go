package main

import "fmt"

type person struct {
	first string
	age   int
}

func main() {
	p := person{
		first: "Hernan",
		age:   53,
	}

	p.speak()
}

func (p person) speak() {
	fmt.Printf("My name is %v and my age is %v\n", p.first, p.age)
}
