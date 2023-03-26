package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Masukkan kalimat: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i := scanner.Text()

	count := make(map[string]int)
	for _, v := range i {
		count[string(v)]++
	}
	fmt.Println("Kalimat input:")
	for j := 0; j < len(i); j++ {
		fmt.Println(string(i[j]))
	}
	fmt.Println(count)
}
