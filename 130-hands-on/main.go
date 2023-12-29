package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	v1 := vehicle{
		engine: engine{true},
		make:   "ford",
		model:  "ka",
		doors:  2,
		color:  "black",
	}

	v2 := vehicle{
		engine: engine{false},
		make:   "ford",
		model:  "taurus",
		doors:  4,
		color:  "white",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.engine.electric)
	fmt.Println(v2.doors)
}
