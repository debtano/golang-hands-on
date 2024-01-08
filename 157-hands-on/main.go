package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.length * s.width
}

func (c Circle) Area() float64 {
	return math.Pow((math.Pi * c.radius), 2)
}

type Shape interface {
	Area() float64
}

func Info(sh Shape) {
	fmt.Println(sh.Area())
}

func main() {

	s1 := Square{
		length: 10.0,
		width:  20.0,
	}

	c1 := Circle{
		radius: 10.0,
	}

	Info(s1)
	Info(c1)
}
