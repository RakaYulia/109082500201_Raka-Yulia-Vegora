package main

import "fmt"

const NMAX = 1000000

// struct partai
type partai struct {
	nama  int
	suara int
}

// tipe tabPartai: array of partai dengan kapasitas NMAX
type tabPartai [NMAX]partai

// posisi: mencari indeks partai dengan nama tertentu, -1 jika tidak ditemukan
func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	n := 0

	// Prompt input
	fmt.Print("Masukkan proses input suara : ")

	// Input suara secara berulang hingga -1
	var input int
	fmt.Scan(&input)
	for input != -1 {
		idx := posisi(p, n, input)
		if idx == -1 {
			p[n].nama = input
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
		fmt.Scan(&input)
	}

	// Insertion sort descending berdasarkan jumlah suara
	// Jika suara sama, urutkan ascending berdasarkan nama partai
	for i := 1; i < n; i++ {
		key := p[i]
		j := i - 1
		for j >= 0 && (p[j].suara < key.suara || (p[j].suara == key.suara && p[j].nama > key.nama)) {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = key
	}

	// Tampilkan hasil
	fmt.Println("\nHasil Perhitungan suara :")
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)
	}
	if n > 0 {
		fmt.Println()
	}
}