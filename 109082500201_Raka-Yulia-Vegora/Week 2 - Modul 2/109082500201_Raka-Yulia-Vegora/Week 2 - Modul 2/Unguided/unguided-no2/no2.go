package main
import"fmt"

func main(){
	berhasil := true
	
	for i := 1; i <= 5; i++ {
		var w1, w2, w3, w4 string
		fmt.Printf("Percobaan %d : ", i)
		fmt.Scan(&w1, &w2, &w3, &w4)
		if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu" {
			berhasil = false
		}
	}
	if berhasil {
		fmt.Println("BERHASIL : true")
		} else {
			fmt.Println("BERHASIL : false")
	}
}