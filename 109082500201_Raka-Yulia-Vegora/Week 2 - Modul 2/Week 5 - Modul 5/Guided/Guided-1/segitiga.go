package main
import"fmt"

// Membuat alias 'desimal' untuk tipe data float64
type desimal float64

func main(){
	// Mendeklrasikan variabel dengan tipe data alias yang baru dibuat
	var alas, tinggi, luas desimal

	fmt.Print("Masukkan panjang alas segitiga: ")
	fmt.Scan(&alas)

	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	// Menghitung luas segitiga
	luas = 0.5 * alas * tinggi

	fmt.Printf("Luas segitiga dengan alas %.2f dan  tinggi  %.2f adalah %.2f\n", alas, tinggi, luas)
}