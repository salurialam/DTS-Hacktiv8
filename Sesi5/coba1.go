package main

import (
	"fmt"
	"os"
)

type teman struct {
	absen     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	dataTeman := []teman{
		{1, "Andi", "Jakarta", "Programmer", "Ingin menguasai bahasa pemrograman baru"},
		{2, "Budi", "Bandung", "Data Analyst", "Meningkatkan skill dalam pengolahan data"},
		{3, "Cici", "Surabaya", "UI/UX Designer", "Menambah wawasan dalam desain aplikasi"},
		{4, "Doni", "Semarang", "Web Developer", "Meningkatkan kemampuan dalam pengembangan web"},
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run coba1.go <nomor_absen>")
		os.Exit(1)
	}

	absen := os.Args[1]

	for _, t := range dataTeman {
		if fmt.Sprintf("%d", t.absen) == absen {
			fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n",
				t.nama, t.alamat, t.pekerjaan, t.alasan)
			return
		}
	}

	fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan")
}
