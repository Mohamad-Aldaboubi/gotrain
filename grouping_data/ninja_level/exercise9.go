package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["mohamad"] = []string{"playing pes", "learning new things", "watching matches"}
	for i, v := range m {
		fmt.Println("record for", i)
		for j, v1 := range v {
			fmt.Println(j, v1)

		}
	}
}
