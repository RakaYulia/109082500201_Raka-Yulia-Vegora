package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, hapus, cari int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nIsi seluruh array:")
	fmt.Println(arr)

	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Print("\nMasukkan nilai x untuk indeks kelipatan x: ")
	fmt.Scan(&x)

	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	fmt.Print("\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&hapus)

	arr = append(arr[:hapus], arr[hapus+1:]...)

	fmt.Println("Array setelah dihapus:")
	fmt.Println(arr)

	total := 0
	for _, v := range arr {
		total += v
	}

	rata := float64(total) / float64(len(arr))
	fmt.Printf("Rata-rata: %.2f\n", rata)

	var jumlahKuadrat float64
	for _, v := range arr {
		jumlahKuadrat += math.Pow(float64(v)-rata, 2)
	}

	sd := math.Sqrt(jumlahKuadrat / float64(len(arr)))
	fmt.Printf("Standar deviasi: %.2f\n", sd)

	fmt.Print("Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)

	frek := 0
	for _, v := range arr {
		if v == cari {
			frek++
		}
	}

	fmt.Println("Frekuensi", cari, "=", frek)
}