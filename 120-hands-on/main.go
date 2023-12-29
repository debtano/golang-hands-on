package main

import "fmt"

func main() {

	jb := []string{"shaken, not stirred", "martinis", "fast cars"}
	mj := []string{"intelligence", "literature", "computer science"}
	nd := []string{"cats", "ice cream", "sunsets"}

	xm := map[string][]string{
		"bond_james":         jb,
		"moneypenney_jenney": mj,
		"no_dr":              nd,
	}

	xm["fleming_iam"] = []string{"steaks", "cigars", "espionage"}

	delete(xm, "no_dr")

	for k, v := range xm {
		for i, va := range v {
			fmt.Println(k, i, va)
		}
	}
}
