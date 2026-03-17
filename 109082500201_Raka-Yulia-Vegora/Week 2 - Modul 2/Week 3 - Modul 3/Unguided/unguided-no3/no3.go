package main
import ("fmt")

// prosedur untuk mencetak deret
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	// cetak suku terakhir
	fmt.Println(1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)

	cetakDeret(n)
}