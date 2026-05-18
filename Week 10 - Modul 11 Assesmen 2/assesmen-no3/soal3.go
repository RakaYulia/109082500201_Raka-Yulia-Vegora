package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")

	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	indeksMax := 0

	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[indeksMax] {
			indeksMax = i
		}
	}

	return indeksMax
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i + 1
		}
	}

	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")

	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := int((1 + tumbuh[i]) * float64(pop[i]))
			fmt.Println(prov[i], prediksi)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var namaCari string

	InputData(&prov, &pop, &tumbuh)

	fmt.Scan(&namaCari)

	indeksCepat := ProvinsiTercepat(tumbuh)
	indeksCari := IndeksProvinsi(prov, namaCari)

	fmt.Println()
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", prov[indeksCepat])
	fmt.Println()
	fmt.Println("Indeks provinsi yang dicari :", indeksCari)
	fmt.Println()

	Prediksi(prov, pop, tumbuh)
}