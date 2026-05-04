package main

import "fmt"

func main() {
	var N int
	var berat [1000]float64

	fmt.Println("=== Program Pendataan Berat Anak Kelinci ===")
	fmt.Print("Masukkan jumlah anak kelinci yang akan ditimbang: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	terkecil := berat[0]
	terbesar := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < terkecil {
			terkecil = berat[i]
		}

		if berat[i] > terbesar {
			terbesar = berat[i]
		}
	}

	fmt.Println()
	fmt.Println("=== Hasil Pendataan Berat Anak Kelinci ===")
	fmt.Printf("Berat anak kelinci terkecil adalah: %.2f kg\n", terkecil)
	fmt.Printf("Berat anak kelinci terbesar adalah: %.2f kg\n", terbesar)
}