package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama   string
	Alamat string
	Kerja  string
	Alasan string
}

var temanData = []Teman{
	{"Ega Nurhaya", "Batam", "Belum Bekerja", "Membutuhkan career coaching dari profesional"},
	{"Alfachri", "Depok", "Streamer", "Menambahkan skill set"},
	{"Ray Fauzi", "Jakarta", "Pelucu", "Mendapatkan ide untuk membuat trailer"},
	{"Luthfi Hamilawan", "Bogor", "Vtuber", "Mempertajam skill programming"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Masukkan sebuah input angka (0, 1, 2, atau 3)")
		return
	}

	index := os.Args[1]
	switch index {
	case "0":
		displayTemanData(temanData[0])
	case "1":
		displayTemanData(temanData[1])
	case "2":
		displayTemanData(temanData[2])
	case "3":
		displayTemanData(temanData[3])
	default:
		fmt.Println("Input tidak sesuai pilih antara 0, 1, 2, atau 3")
	}
}

func displayTemanData(teman Teman) {
	fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n", teman.Nama, teman.Alamat, teman.Kerja, teman.Alasan)
}
