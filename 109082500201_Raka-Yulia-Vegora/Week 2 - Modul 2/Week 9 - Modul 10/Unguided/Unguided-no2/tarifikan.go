package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64

	fmt.Println("=== Program Pendataan Berat Ikan ===")
	fmt.Print("Masukkan jumlah ikan dan kapasitas tiap wadah: ")
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	totalSemua := 0.0
	jumlahWadah := 0

	fmt.Println("\n=== Total Berat Ikan Setiap Wadah ===")
	for i := 0; i < x; i += y {
		totalWadah := 0.0

		for j := i; j < i+y && j < x; j++ {
			totalWadah += berat[j]
		}

		jumlahWadah++
		totalSemua += totalWadah
		fmt.Printf("Wadah ke-%d: %.2f kg\n", jumlahWadah, totalWadah)
	}

	rataRata := totalSemua / float64(jumlahWadah)

	fmt.Println("\n=== Hasil Akhir ===")
	fmt.Printf("Rata-rata berat ikan setiap wadah: %.2f kg\n", rataRata)
}