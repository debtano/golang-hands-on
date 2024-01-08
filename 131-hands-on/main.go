package main

import "fmt"

func main() {

	s1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "hernan",
		friends: map[string]int{
			"jose":  1,
			"pedro": 2,
		},
		favDrinks: []string{"whisky", "in", "the", "jar"},
	}

	for k, v := range s1.friends {
		fmt.Println(k, v)
	}

	for _, v := range s1.favDrinks {
		fmt.Println(v)
	}
}
