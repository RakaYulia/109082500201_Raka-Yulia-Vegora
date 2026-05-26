package main

import "fmt"

func main() {
	arr := make([]int, 0)

	// Baca input sampai bilangan negatif
	for {
		var x int
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		arr = append(arr, x)
	}

	// Insertion Sort ascending
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	// Cetak array terurut
	for i := 0; i < len(arr); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
	fmt.Println()

	// Cek jarak
	tetap := true
	jarak := 0
	if len(arr) >= 2 {
		jarak = arr[1] - arr[0]
		for i := 2; i < len(arr); i++ {
			if arr[i]-arr[i-1] != jarak {
				tetap = false
				break
			}
		}
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}