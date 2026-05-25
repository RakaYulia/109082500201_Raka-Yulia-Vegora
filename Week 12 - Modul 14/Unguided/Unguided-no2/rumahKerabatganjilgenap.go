package main

import "fmt"

func bubbleSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func bubbleSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var count int
		fmt.Scan(&count)

		odds := []int{}
		evens := []int{}

		for j := 0; j < count; j++ {
			var num int
			fmt.Scan(&num)
			if num%2 != 0 {
				odds = append(odds, num)
			} else {
				evens = append(evens, num)
			}
		}

		// Ganjil urut naik
		bubbleSortAsc(odds)

		// Genap urut turun
		bubbleSortDesc(evens)

		// Cetak dengan label Daerah
		fmt.Printf("Daerah %d:", i+1)
		for _, v := range odds {
			fmt.Printf(" %d", v)
		}
		for _, v := range evens {
			fmt.Printf(" %d", v)
		}
		fmt.Println()
	}
}