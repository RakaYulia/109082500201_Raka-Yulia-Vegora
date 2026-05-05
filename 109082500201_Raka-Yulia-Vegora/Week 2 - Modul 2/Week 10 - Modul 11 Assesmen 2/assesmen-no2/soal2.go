package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   int
	Nama  string
	Nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(N)

	if *N > nMax {
		*N = nMax
	}

	for i := 0; i < *N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].Nama, &T[i].Nilai)
	}
}

func nilaiPertama(T arrayMahasiswa, N int, nim int) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			return T[i].Nilai
		}
	}

	return -1
}

func nilaiTerbesar(T arrayMahasiswa, N int, nim int) int {
	max := -1

	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			if T[i].Nilai > max {
				max = T[i].Nilai
			}
		}
	}

	return max
}

func tampilkanHasil(T arrayMahasiswa, N int, nim int) {
	pertama := nilaiPertama(T, N, nim)
	terbesar := nilaiTerbesar(T, N, nim)

	if pertama == -1 {
		fmt.Println("NIM tidak ditemukan")
	} else {
		fmt.Printf("Nilai pertama dari NIM %d adalah %d\n", nim, pertama)
		fmt.Printf("Nilai Terbesar dari NIM %d adalah %d\n", nim, terbesar)
	}
}

func main() {
	var data arrayMahasiswa
	var N int
	var cariNIM int

	inputData(&data, &N)

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&cariNIM)

	tampilkanHasil(data, N, cariNIM)
}
