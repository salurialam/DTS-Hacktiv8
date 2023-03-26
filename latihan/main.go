package main

import (
	"fmt"
)

func main() {
	// Input kalimat
	var kalimat string
	fmt.Print("Masukkan kalimat: ")
	fmt.Scanln(&kalimat)

	// Cetak kalimat
	fmt.Println("Kalimat yang Anda masukkan adalah:", kalimat)
}
