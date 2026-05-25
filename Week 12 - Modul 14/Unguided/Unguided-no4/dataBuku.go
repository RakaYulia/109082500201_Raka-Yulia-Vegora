package main
import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [1 + nMax]Buku

var Pustaka DaftarBuku
var nPustaka int

func cetakGaris() {
	fmt.Println("================================================")
}

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 1; i <= *n; i++ {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	terfavorit := pustaka[1]
	for i := 2; i <= n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}
	cetakGaris()
	fmt.Println("           ★  BUKU TERFAVORIT  ★")
	cetakGaris()
	fmt.Printf("  ID       : %s\n", terfavorit.id)
	fmt.Printf("  Judul    : %s\n", terfavorit.judul)
	fmt.Printf("  Penulis  : %s\n", terfavorit.penulis)
	fmt.Printf("  Penerbit : %s\n", terfavorit.penerbit)
	fmt.Printf("  Tahun    : %d\n", terfavorit.tahun)
	fmt.Printf("  Rating   : %d/10\n", terfavorit.rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := n
	if batas > 5 {
		batas = 5
	}
	cetakGaris()
	fmt.Println("        ★  5 BUKU RATING TERTINGGI  ★")
	cetakGaris()
	for i := 1; i <= batas; i++ {
		fmt.Printf("  %d. %-20s\n", i, pustaka[i].judul)
		fmt.Printf("     Penulis  : %s\n", pustaka[i].penulis)
		fmt.Printf("     Penerbit : %s\n", pustaka[i].penerbit)
		fmt.Printf("     Rating   : %d/10\n", pustaka[i].rating)
		if i < batas {
			fmt.Println("  ------------------------------------------------")
		}
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	lo, hi := 1, n
	idx := -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if pustaka[mid].rating == r {
			idx = mid
			break
		} else if pustaka[mid].rating > r {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	cetakGaris()
	fmt.Printf("      ★  HASIL PENCARIAN RATING %d  ★\n", r)
	cetakGaris()
	if idx == -1 {
		fmt.Printf("  [!] Tidak ada buku dengan rating %d\n", r)
	} else {
		fmt.Printf("  ID        : %s\n", pustaka[idx].id)
		fmt.Printf("  Judul     : %s\n", pustaka[idx].judul)
		fmt.Printf("  Penulis   : %s\n", pustaka[idx].penulis)
		fmt.Printf("  Penerbit  : %s\n", pustaka[idx].penerbit)
		fmt.Printf("  Tahun     : %d\n", pustaka[idx].tahun)
		fmt.Printf("  Eksemplar : %d\n", pustaka[idx].eksemplar)
		fmt.Printf("  Rating    : %d/10\n", pustaka[idx].rating)
	}
	cetakGaris()
}

func main() {
	DaftarkanBuku(&Pustaka, &nPustaka)
	CetakTerfavorit(Pustaka, nPustaka)
	UrutBuku(&Pustaka, nPustaka)
	Cetak5Terbaru(Pustaka, nPustaka)

	var r int
	fmt.Scan(&r)
	CariBuku(Pustaka, nPustaka, r)
}