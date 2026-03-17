package main
import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var x, y, z int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&y)

	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&z)

	// hasil komposisi
	result1 := f(g(h(x))) // F(G(H(x)))
	result2 := g(h(f(y))) // G(H(F(y)))
	result3 := h(f(g(z))) // H(F(G(z)))

	fmt.Printf("\nF(G(H(%d))) : %d\n", x, result1)
	fmt.Printf("G(H(F(%d))) : %d\n", y, result2)
	fmt.Printf("H(F(G(%d))) : %d\n", z, result3)
}