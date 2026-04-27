package main
import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var s string
	*n = 0

	fmt.Print("Teks : ")
	for *n < NMAX {
		fmt.Scan(&s)

		if s == "." {
			break
		}

		(*t)[*n] = []rune(s)[0]
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var balik tabel

	for i := 0; i < n; i++ {
		balik[i] = t[i]
	}

	balikanArray(&balik, n)

	for i := 0; i < n; i++ {
		if t[i] != balik[i] {
			return false
		}
	}

	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom : true")
	} else {
		fmt.Println("Palindrom : false")
	}

	balikanArray(&tab, n)

	fmt.Print("Reverse : ")
	cetakArray(tab, n)
}