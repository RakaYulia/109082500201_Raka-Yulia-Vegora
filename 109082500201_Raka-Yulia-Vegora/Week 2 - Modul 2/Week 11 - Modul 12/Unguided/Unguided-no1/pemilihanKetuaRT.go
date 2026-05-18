package main

import "fmt"

func main() {
	votes := make(map[int]int)
	count := 0
	valid := 0

	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		count++
		if n >= 1 && n <= 20 {
			valid++
			votes[n]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", count)
	fmt.Printf("Suara sah: %d\n", valid)

	for i := 1; i <= 20; i++ {
		if votes[i] > 0 {
			fmt.Printf("%d: %d\n", i, votes[i])
		}
	}
}