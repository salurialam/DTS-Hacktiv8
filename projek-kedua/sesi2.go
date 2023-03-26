package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ", i)
	}

	keys := []string{"U+0421", "U+0410", "U+0428", "U+0410", "U+0420", "U+0412", "U+041E"}
	values := []string{"C", "A", "Ш", "A", "P", "B", "O"}

	unicode := make(map[string]string)

	for j := range keys {
		unicode[keys[j]] = values[j]
	}
	p := 0
	for j := 0; j <= 10; j++ {
		if j == 5 {
			for _, k := range keys {
				fmt.Printf("character %#v %s starts at byte positions %d \n", k, unicode[k], p)
				p = p + 2
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}
}
