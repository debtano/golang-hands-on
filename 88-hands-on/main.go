package main

import "fmt"

func main() {
	m := map[string]int{
		"Hernan": 53,
		"Marisa": 59,
	}
	for k, v := range m {
		fmt.Printf("Key is %v\t\t value is %v\n", k, v)
	}

	age := m["Hernan"]
	fmt.Printf("Value of Hernan is : %v\n", age)

	q := m["Q"]
	fmt.Printf("Value of Q is : %v\n", q)

	if value, ok := m["Q"]; ok {
		fmt.Println(value)
	}
	fmt.Println("No value for Q")
}
