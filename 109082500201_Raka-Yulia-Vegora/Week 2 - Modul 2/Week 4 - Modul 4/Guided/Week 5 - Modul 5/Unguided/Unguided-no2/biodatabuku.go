package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// tipe data alias
type angka int
type kata string

// struct Buku
type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

// fungsi input string
func inputKata(reader *bufio.Reader, pesan string) kata {
	fmt.Print(pesan)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return kata(text)
}

// fungsi input integer
func inputAngka(reader *bufio.Reader, pesan string) angka {
	fmt.Print(pesan)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	nilai, _ := strconv.Atoi(text)
	return angka(nilai)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buku Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	buku.judul = inputKata(reader, "Masukkan judul buku : ")
	buku.penulis = inputKata(reader, "Masukkan nama penulis : ")
	buku.penerbit = inputKata(reader, "Masukkan penerbit : ")
	buku.tahunTerbit = inputAngka(reader, "Masukkan tahun terbit : ")
	buku.jumlahHalaman = inputAngka(reader, "Masukkan jumlah halaman: ")

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Println("Judul Buku    :", buku.judul)
	fmt.Println("Penulis       :", buku.penulis)
	fmt.Println("Penerbit      :", buku.penerbit)
	fmt.Println("Tahun Terbit  :", buku.tahunTerbit)
	fmt.Println("Jumlah Halaman:", buku.jumlahHalaman)
}