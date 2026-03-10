package main
import "fmt"

func main(){
	var tahun int 
	fmt.Print("Masukkan Tahun: ") 
	fmt.Scan(&tahun) 
	
	if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0){
		fmt.Println("Kabisat : true")
		} else {
			fmt.Println("kabisat : false")
	}
}
