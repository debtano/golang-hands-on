package main

import "fmt"

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenney", "I am 008"}

	xp := [][]string{jb, jm}
	for i, v := range xp {
		fmt.Println(i, v)
	}
}
