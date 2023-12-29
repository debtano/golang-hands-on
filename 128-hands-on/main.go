package main

import "fmt"

type person struct {
	first   string
	last    string
	fav_ice []string
}

func main() {

	p1 := person{
		first:   "hernan",
		last:    "antolini",
		fav_ice: []string{"dulce", "vainilla"},
	}

	p2 := person{
		first:   "marisa",
		last:    "zamora",
		fav_ice: []string{"choc", "zap"},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.fav_ice {
		fmt.Println(v)
	}

	fmt.Println(p2.first, p2.last)
	for _, v := range p2.fav_ice {
		fmt.Println(v)
	}

	mp := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println("-------------------------")

	for _, v := range mp {
		fmt.Println(v.first, v.last)
		for _, y := range v.fav_ice {
			fmt.Println(y)
		}

	}

}
