package main
import"fmt"

func main(){
	var berat, kg, gram, biayaKg, bGram int
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&berat)
	
	kg, gram = berat/1000, berat%1000
	biayaKg = kg * 10000

// aturan biaya sisa berat
	if kg > 10 {
		bGram = 0
		} else if gram >= 500 {
			bGram = gram * 5
			} else {
				bGram = gram * 15
			}
			fmt.Println("\n===== Detail Perhitungan =====")
			fmt.Printf("Detail berat : %2d kg + %3d gram\n", kg, gram)
			fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, bGram)
			fmt.Printf("Total biaya: Rp %d\n", biayaKg+bGram)
}

