package main
import"fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}
func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}
func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
func main() {
	var a, b, c, d int
	fmt.Print("Masukan nilai a : ")
	fmt.Scanln(&a)
	fmt.Print("Masukan nilai b : ")
	fmt.Scanln(&b)
	fmt.Print("Masukan nilai c : ")
	fmt.Scanln(&c)
	fmt.Print("Masukan nilai d : ")
	fmt.Scanln(&d)
	
	fmt.Printf("Hasil permutasi %d dan %d adalah : %d\n", a, c, permutasi(a, c))
	fmt.Printf("Hasil kombinasi %d dan %d adalah : %d\n", a, c, kombinasi(a, c))
	fmt.Printf("Hasil permutasi %d dan %d adalah : %d\n", b, d, permutasi(b, d))
	fmt.Printf("Hasil kombinasi %d dan %d adalah : %d\n\n", b, d, kombinasi(b, d))
}