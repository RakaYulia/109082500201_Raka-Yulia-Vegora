package main
import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else { // mancanegara
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
}

func biayaPerHari(jumlahMhs int) int {
	// biaya domestic per hari per mahasiswa:
	// makan (2x) = 2 x 35.000 = 70.000
	// penginapan = 250.000
	// uang saku = 300.000
	// total = 620.000
	return 620000 * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	biayaHarian := biayaPerHari(jumlahMhs)

	if tujuan == "mancanegara" {
		biayaHarian = int(float64(biayaHarian) * 1.5)
	}

	*totalBiaya = float64(hariDitanggung * biayaHarian)
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}